package tuple

type Pair[T, U any] interface {
	Values() (T, U)
	sealed()
}

type Triple[T, U, V any] interface {
	Values() (T, U, V)
	sealed()
}

type Quadruple[T, U, V, W any] interface {
	Values() (T, U, V, W)
	sealed()
}

func NewPair[T, U any](t T, u U) Pair[T, U] {
	return pair[T, U]{first: t, second: u}
}

func NewTriple[T, U, V any](t T, u U, v V) Triple[T, U, V] {
	return triple[T, U, V]{first: t, second: u, third: v}
}

func NewQuadruple[T, U, V, W any](t T, u U, v V, w W) Quadruple[T, U, V, W] {
	return quadruple[T, U, V, W]{first: t, second: u, third: v, fourth: w}
}

type pair[T, U any] struct {
	first  T
	second U
}

func (pair[T, U]) sealed() {}

func (p pair[T, U]) Values() (T, U) {
	return p.first, p.second
}

type triple[T, U, V any] struct {
	first  T
	second U
	third  V
}

func (triple[T, U, V]) sealed() {}

func (t triple[T, U, V]) Values() (T, U, V) {
	return t.first, t.second, t.third
}

type quadruple[T, U, V, W any] struct {
	first  T
	second U
	third  V
	fourth W
}

func (quadruple[T, U, V, W]) sealed() {}

func (q quadruple[T, U, V, W]) Values() (T, U, V, W) {
	return q.first, q.second, q.third, q.fourth
}
