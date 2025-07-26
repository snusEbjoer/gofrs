package seq

import (
	"iter"
)

func FromSlice[T any](arr []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range arr {
			if !yield(v) {
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

	return func(yield func(R, []T) bool) {
		for k, v := range res {
			if !yield(k, v) {
				break
			}
		}
	}
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

func Merge[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			for v := range seq {
				if !yield(v) {
					break
				}
			}
		}
	}
}

type puller[T any] struct {
	stop  func()
	next  func() (T, bool)
	ended bool
}

func Interleave[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		pullers := make([]puller[T], 0, len(seqs))

		for _, seq := range seqs {
			next, stop := iter.Pull(seq)

			pullers = append(pullers, puller[T]{
				next: next,
				stop: stop,
			})
		}

		seqsEnded := 0

		for seqsEnded < len(pullers) {
			for i, itr := range pullers {
				if !itr.ended {
					if v, ok := itr.next(); ok {
						if !yield(v) {
							return
						}
					} else {
						itr.stop()
						pullers[i].ended = true
						seqsEnded++
					}
				}
			}
		}
	}
}

func Enumerate[T any](seq iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for v := range seq {
			if !yield(i, v) {
				break
			}

			i++
		}
	}
}

func Zip[T, R any](seq1 iter.Seq[T], seq2 iter.Seq[R]) iter.Seq2[T, R] {
	return func(yield func(T, R) bool) {
		next1, stop1 := iter.Pull(seq1)
		next2, stop2 := iter.Pull(seq2)
		defer stop1()
		defer stop2()

		var ended bool

		for !ended {
			v1, ok1 := next1()
			v2, ok2 := next2()
			if ok1 && ok2 {
				if !yield(v1, v2) {
					return
				}
			} else {
				ended = true
			}
		}
	}
}
