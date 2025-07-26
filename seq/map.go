package seq

import "iter"

func Map[T, R any](seq iter.Seq[T], f func(T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range seq {
			if !yield(f(v)) {
				break
			}
		}
	}
}

func MapToSeq2[T, R1, R2 any](seq iter.Seq[T], f func(T) (R1, R2)) iter.Seq2[R1, R2] {
	return func(yield func(R1, R2) bool) {
		for v := range seq {
			if !yield(f(v)) {
				break
			}
		}
	}
}

func FilterMapToSeq2[T, R1, R2 any](seq iter.Seq[T], f func(T) (R1, R2, bool)) iter.Seq2[R1, R2] {
	return func(yield func(R1, R2) bool) {
		for v := range seq {
			if v1, v2, ok := f(v); ok {
				if !yield(v1, v2) {
					break
				}
			}
		}
	}
}

func FilterMap[T, R any](seq iter.Seq[T], f func(T) (R, bool)) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range seq {
			if v1, ok := f(v); ok {
				if !yield(v1) {
					break
				}
			}

		}
	}
}

func DistinctMap[T any, R comparable](seq iter.Seq[T], f func(T) R) iter.Seq[R] {
	return Distinct(Map(seq, f))
}
