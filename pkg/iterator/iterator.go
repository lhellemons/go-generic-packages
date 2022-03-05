package iterator

import (
	"generic-packages/pkg/list"
	"generic-packages/pkg/predicate"
)

type Iterator[T any] interface {
	Next() (t T, ok bool)
	sealed()
}

type iterator[T any] struct {
	done bool
	next func() (T, bool)
}

func (iterator[T]) sealed() {}

func (i *iterator[T]) Next() (t T, ok bool) {
	if i.done {
		return t, false
	}
	t, ok = i.next()
	if !ok {
		i.done = true
		return t, false
	}
	return t, true
}

// constructors

// Callback creates an Iterator that gets each Next value by calling a callback func.
// The func should return the next value and true if there is a next value, and ok = false if there is not.
// If ok = false, the value of t will be discarded, and subsequent calls to Next will not invoke the callback anymore.
func Callback[T any](f func() (t T, ok bool)) Iterator[T] {
	return &iterator[T]{next: f}
}

func Empty[T any]() Iterator[T] {
	return Callback(func() (t T, ok bool) { return t, ok })
}

func List[T any](l list.List[T]) Iterator[T] {
	return Slice(l.Items())
}

func Slice[T any](ts []T) Iterator[T] {
	i := 0
	return Callback(func() (t T, ok bool) {
		if i == len(ts) {
			return t, false
		}
		i++
		return ts[i-1], true
	})
}

func Static[T any](ts ...T) Iterator[T] {
	return Slice(ts)
}

func Chan[T any](c <-chan T) Iterator[T] {
	return Callback(func() (T, bool) {
		t, ok := <-c
		return t, ok
	})
}

func Naturals() Iterator[int] {
	i := -1
	return Callback(func() (int, bool) {
		i++
		return i, true
	})
}

// map

func Map[T, U any](f func(T) U, it Iterator[T]) Iterator[U] {
	return Callback(func() (u U, ok bool) {
		if t, ok := it.Next(); ok {
			return f(t), true
		}
		return u, false
	})
}

func FlatMap[T, U any](f func(T) Iterator[U], it Iterator[T]) Iterator[U] {
	return Flatten(Map(f, it))
}

func Flatten[T any](it Iterator[Iterator[T]]) Iterator[T] {
	var curr = Empty[T]()
	var currOk bool
	return Callback(func() (t T, ok bool) {
		if t, ok := curr.Next(); ok {
			return t, ok
		}
		curr, currOk = it.Next()
		if currOk {
			return curr.Next()
		}
		return t, false
	})
}

func ForEach[T any](callback func(T), it Iterator[T]) {
	done := false
	for !done {
		t, ok := it.Next()
		if !ok {
			done = true
			break
		}
		callback(t)
	}
}

func Stream[T any](it Iterator[T]) <-chan T {
	c := make(chan T)
	go ForEach(func(t T) { c <- t }, it)
	return c
}

func Collect[T any](it Iterator[T]) (ts []T) {
	ForEach(func(t T) { ts = append(ts, t) }, it)
	return ts
}

func Filter[T any](f func(T) bool, it Iterator[T]) Iterator[T] {
	return Callback(func() (t T, ok bool) {
		for {
			t, ok = it.Next()
			if !ok {
				return
			}

			if f(t) {
				return
			}
		}
	})
}

func TakeWhile[T any](f func(T) bool, it Iterator[T]) Iterator[T] {
	done := false
	return Callback(func() (t T, ok bool) {
		if done {
			return t, false
		}

		t, ok = it.Next()

		if !ok {
			done = true
			return
		}
		if !f(t) {
			done = true
			return
		}
		return
	})
}

func TakeUntil[T any](pred func(T) bool, it Iterator[T]) Iterator[T] {
	return TakeWhile(predicate.Not(pred), it)
}

func DropWhile[T any](pred func(T) bool, it Iterator[T]) Iterator[T] {
	satisfied := false
	return Callback(func() (t T, ok bool) {
		for !satisfied {
			t, ok = it.Next()
			if !ok {
				return
			}
			if !pred(t) {
				satisfied = true
				return t, ok
			}
		}
		return it.Next()
	})
}

func DropUntil[T any](pred func(T) bool, it Iterator[T]) Iterator[T] {
	return DropWhile(predicate.Not(pred), it)
}
