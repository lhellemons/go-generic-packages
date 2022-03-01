package list

type List[T any] interface {
	sealed()
	Items() []T
}

type list[T any] []T

func (l list[T]) sealed() {}

func (l list[T]) Items() []T {
	ts := make([]T, len(l))
	copy(ts, l)
	return ts
}

func Empty[T any]() List[T] {
	return list[T]{}
}

func New[T any](ts ...T) List[T] {
	return FromSlice(ts)
}

func FromSlice[T any](ts []T) List[T] {
	l := make(list[T], len(ts))
	copy(l, ts)
	return l
}

// === Map

type Mapper[T, U any] interface {
	Map(T) U
}

type MapperFunc[T, U any] func(T) U

func (f MapperFunc[T, U]) Map(t T) U {
	return f(t)
}

func Map[T, U any](f func(T) U, ts List[T]) List[U] {
	tsItems := ts.Items()
	us := make(list[U], len(tsItems))
	for i, t := range tsItems {
		us[i] = f(t)
	}
	return us
}
