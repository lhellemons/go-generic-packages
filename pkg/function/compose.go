package function

func Compose[T, U, V any](u func(T) U, v func(U) V) func(T) V {
	return func(t T) V {
		return v(u(t))
	}
}

func Compose3[T, U, V, W any](u func(T) U, v func(U) V, w func(V) W) func(T) W {
	return func(t T) W {
		return w(v(u(t)))
	}
}

func Compose4[T, U, V, W, X any](u func(T) U, v func(U) V, w func(V) W, x func(W) X) func(T) X {
	return func(t T) X {
		return x(w(v(u(t))))
	}
}

func Compose5[T, U, V, W, X, Y any](u func(T) U, v func(U) V, w func(V) W, x func(W) X, y func(X) Y) func(T) Y {
	return func(t T) Y {
		return y(x(w(v(u(t)))))
	}
}

func Compose6[T, U, V, W, X, Y, Z any](u func(T) U, v func(U) V, w func(V) W, x func(W) X, y func(X) Y, z func(Y) Z) func(T) Z {
	return func(t T) Z {
		return z(y(x(w(v(u(t))))))
	}
}

func Wrap[T, U any](f func(T) U, w func(T, func(T) U) U) func(T) U {
	return func(t T) U {
		return w(t, f)
	}
}
