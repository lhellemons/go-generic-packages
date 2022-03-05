package predicate

// TODO add tests

type Predicate[T any] func(T) bool

func Not[T any](f func(T) bool) func(T) bool {
	return func(t T) bool {
		return !f(t)
	}
}

func And[T any](f func(T) bool, gs ...func(T) bool) func(T) bool {
	fs := static(f, gs...)
	return func(t T) bool {
		for _, f := range fs {
			if !f(t) {
				return false
			}
		}
		return true
	}
}

func Or[T any](f func(T) bool, gs ...func(T) bool) func(T) bool {
	fs := static(f, gs...)
	return func(t T) bool {
		for _, f := range fs {
			if f(t) {
				return true
			}
		}
		return false
	}
}

func static[T any](t T, ts ...T) (tss []T) {
	tss = append(tss, t)
	tss = append(tss, ts...)
	return tss
}
