package gofrs

import "iter"

func FromSlice[T any](arr []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range arr {
			if !yield(v) {
				break
			}
		}
	}
}

func FromMap[K comparable, V any](mp map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range mp {
			if !yield(k, v) {
				break
			}
		}
	}
}

func Map[T, R any](seq iter.Seq[T], f func(T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range seq {
			if !yield(f(v)) {
				break
			}
		}
	}
}

func Filter[T any](seq iter.Seq[T], pred func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if pred(v) {
				if !yield(v) {
					break
				}
			}
		}
	}
}

func FilterMap[T, R any](seq iter.Seq[T], pred func(T) bool, f func(T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range seq {
			if pred(v) {
				if !yield(f(v)) {
					break
				}
			}
		}
	}
}

func Reduce[T, R any](seq iter.Seq[T], f func(T, R) R, initialAcc ...R) R {
	var acc R
	if len(initialAcc) >= 1 {
		acc = initialAcc[0]
	}

	for v := range seq {
		acc = f(v, acc)
	}

	return acc
}

func Distinct[T comparable](seq iter.Seq[T]) iter.Seq[T] {
	seen := make(map[T]struct{})
	res := []T{}

	for v := range seq {
		if _, ok := seen[v]; !ok {
			res = append(res, v)
		} else {
			seen[v] = struct{}{}
		}
	}

	return FromSlice(res)
}

func DistinctMap[T any, R comparable](seq iter.Seq[T], f func(T) R) iter.Seq[R] {
	return Distinct(Map(seq, f))
}

func DistinctBy[T any, R comparable](seq iter.Seq[T], f func(T) R) iter.Seq[T] {
	seen := make(map[R]struct{})
	res := []T{}

	for v := range seq {
		if _, ok := seen[f(v)]; !ok {
			res = append(res, v)
		} else {
			seen[f(v)] = struct{}{}
		}
	}

	return FromSlice(res)
}

func GroupBy[T any, R comparable](seq iter.Seq[T], f func(T) R) iter.Seq2[R, []T] {
	res := make(map[R][]T)

	for v := range seq {
		res[f(v)] = append(res[f(v)], v)
	}

	return FromMap(res)
}

func Take[T any](seq iter.Seq[T], count int) iter.Seq[T] {
	return func(yield func(T) bool) {
		var i int
		for v := range seq {
			if i >= count {
				break
			}
			i++

			if !yield(v) {
				break
			}
		}
	}
}

func Drop[T any](seq iter.Seq[T], count int) iter.Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		for v := range seq {
			if i < count {
				i++
				continue
			}

			if !yield(v) {
				break
			}
		}
	}
}
