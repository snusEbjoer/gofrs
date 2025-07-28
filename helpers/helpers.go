package helpers

type Predicate[T any] = func(v T) bool
type F1[T, R any] = func(T) R
type F2[T, T1, R any] = func(T, T1) R
type Lazy[T any] = func() T

func Identity[T any](v T) T {
	return v
}

func Const1[T, R any](f func() R) F1[T, R] {
	return func(t T) R {
		return f()
	}
}

func NotNil[T any](t *T) bool {
	return t != nil
}
