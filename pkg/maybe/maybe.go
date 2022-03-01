package maybe

type Maybe[T any] interface {
	Val() (T, bool)
	sealed()
}

type maybe[T any] struct {
	hasVal bool
	val    T
}

func (m maybe[T]) sealed() {}

func (m maybe[T]) Val() (T, bool) {
	return m.val, m.hasVal
}

// === constructors

func Just[T any](val T) Maybe[T] {
	return maybe[T]{hasVal: true, val: val}
}

func Nothing[T any]() Maybe[T] {
	return maybe[T]{hasVal: false}
}

// === Map

func Map[T, U any](f func(T) U, m Maybe[T]) Maybe[U] {
	if v, ok := m.Val(); ok {
		return Just(f(v))
	}
	return Nothing[U]()
}

// === FlatMap

func FlatMap[T, U any](f func(T) Maybe[U], m Maybe[T]) Maybe[U] {
	if v, ok := m.Val(); ok {
		return f(v)
	}
	return Nothing[U]()
}
