package helpers

type Predicate[T any] = func(v T) bool
type F1[T, R any] = func(T) R
type F2[T, T1, R any] = func(T, T1) R

func Identity[T any](v T) T {
	return v
}
