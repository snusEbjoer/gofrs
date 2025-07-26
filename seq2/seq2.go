package seq2

import "iter"

func FromMap[K comparable, V any](mp map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range mp {
			if !yield(k, v) {
				break
			}
		}
	}
}

func FromKeys[K comparable, V any](mp map[K]V) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range mp {
			if !yield(k) {
				break
			}
		}
	}
}

func FromValues[K comparable, V any](mp map[K]V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range mp {
			if !yield(v) {
				break
			}
		}
	}
}

func Map[T1, T2, R1, R2 any](seq iter.Seq2[T1, T2], f func(T1, T2) (R1, R2)) iter.Seq2[R1, R2] {
	return func(yield func(R1, R2) bool) {
		for v1, v2 := range seq {
			if !yield(f(v1, v2)) {
				break
			}
		}
	}
}

func MapToSeq1[T1, T2, R any](seq iter.Seq2[T1, T2], f func(T1, T2) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v1, v2 := range seq {
			if !yield(f(v1, v2)) {
				break
			}
		}
	}
}

func FilterMap[T1, T2, R1, R2 any](seq iter.Seq2[T1, T2], f func(T1, T2) (R1, R2, bool)) iter.Seq2[R1, R2] {
	return func(yield func(R1, R2) bool) {
		for v1, v2 := range seq {
			val1, val2, ok := f(v1, v2)
			if ok {
				if !yield(val1, val2) {
					break
				}
			}
		}
	}
}

func FilterMapToSeq1[T1, T2, R any](seq iter.Seq2[T1, T2], f func(T1, T2) (R, bool)) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v1, v2 := range seq {
			v, ok := f(v1, v2)
			if ok {
				if !yield(v) {
					break
				}
			}
		}
	}
}

func CollectErr[T any](seq iter.Seq2[T, error]) ([]T, error) {
	res := []T{}

	for v, err := range seq {
		if err != nil {
			return nil, err
		}

		res = append(res, v)
	}

	return res, nil
}

func CollectOk[T any](seq iter.Seq2[T, bool]) ([]T, bool) {
	res := []T{}

	for v, ok := range seq {
		if !ok {
			return nil, false
		}

		res = append(res, v)
	}

	return res, true
}
