package result

type Result[T any] interface {
	Result() (T, error)
	sealed()
}

type result[T any] struct {
	result T
	error  error
}

func (r result[T]) sealed() {}

func (r result[T]) Result() (T, error) {
	return r.result, r.error
}

func OK[T any](t T) Result[T] {
	return result[T]{result: t}
}

func Error[T any](err error) Result[T] {
	return result[T]{error: err}
}

func Map[T, U any](f func(T) U, r Result[T]) Result[U] {
	t, err := r.Result()
	if err != nil {
		return Error[U](err)
	}
	return OK(f(t))
}

func Flatmap[T, U any](f func(T) Result[U], r Result[T]) Result[U] {
	t, err := r.Result()
	if err != nil {
		return Error[U](err)
	}

	return f(t)
}
