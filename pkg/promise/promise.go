package promise

import (
	"context"
	"errors"
	"sync"

	"generic-packages/pkg/result"
)

var CancelledErr = errors.New("promise cancelled")

type Promise[T any] interface {
	Await() result.Result[T]
	sealed()
}

type promise[T any] struct {
	hasResult bool
	result    result.Result[T]
	getResult func() result.Result[T]
	mx        sync.Mutex
}

func (p *promise[T]) sealed() {}

func (p *promise[T]) Await() result.Result[T] {
	if p.hasResult {
		return p.result
	}

	p.mx.Lock()
	defer p.mx.Unlock()

	p.result = p.getResult()

	p.hasResult = true

	return p.result
}

// constructors

func Lazy[T any](f func() result.Result[T]) Promise[T] {
	return &promise[T]{getResult: f}
}

func Eager[T any](f func() result.Result[T]) Promise[T] {
	p := Lazy(f)
	p.Await()
	return p
}

func Resolved[T any](r result.Result[T]) Promise[T] {
	return &promise[T]{hasResult: true, result: r}
}

func Cancellable[T any](ctx context.Context, f func(ctx context.Context) result.Result[T]) Promise[T] {

	fetch := func() <-chan result.Result[T] {
		c := make(chan result.Result[T])
		go func(ctx context.Context, c chan<- result.Result[T]) {
			r := f(ctx)
			c <- r
			close(c)
		}(ctx, c)

		return c
	}

	return Lazy(func() result.Result[T] {
		select {
		case r := <-fetch():
			return r
		case <-ctx.Done():
			return result.Error[T](CancelledErr)
		}
	})
}

// Map

func Map[T, U any](f func(T) U, p Promise[T]) Promise[U] {
	return Lazy(func() result.Result[U] {
		return result.Map(f, p.Await())
	})
}

// Flatmap

func Flatmap[T, U any](f func(T) Promise[U], p Promise[T]) Promise[U] {
	return Flatten(Map(f, p))
}

func Flatten[T any](p Promise[Promise[T]]) Promise[T] {
	r := p.Await()
	t, err := r.Result()
	if err != nil {
		return Resolved(result.Error[T](err))
	}
	return t
}
