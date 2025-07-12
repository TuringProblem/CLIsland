package utils

func Compose[T, U, V any](f func(U) V, g func(T) U) func(T) V {
	return func(t T) V {
		return f(g(t))
	}
}
