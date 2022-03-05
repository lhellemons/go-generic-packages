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

func Static[T any](ts ...T) List[T] {
	return Slice(ts)
}

func Slice[T any](ts []T) List[T] {
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

func Flatten[T any](tts List[List[T]]) List[T] {
	var ts []T
	for _, t := range tts.Items() {
		ts = append(ts, t.Items()...)
	}
	return Slice(ts)
}

func FlatMap[T, U any](f func(T) List[U], ts List[T]) List[U] {
	return Flatten(Map(f, ts))
}
