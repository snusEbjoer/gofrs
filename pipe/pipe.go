package pipe

import (
	"gofrs"
	"iter"
)

func Pipe[T any](seq iter.Seq[T], fns ...func(iter.Seq[T]) iter.Seq[T]) iter.Seq[T] {
	for _, fn := range fns {
		seq = fn(seq)
	}

	return seq
}

func Map[T, R any](f func(T) R) func(iter.Seq[T]) iter.Seq[R] {
	return func(seq iter.Seq[T]) iter.Seq[R] {
		return gofrs.Map(seq, f)
	}
}

func Filter[T any](pred func(T) bool) func(iter.Seq[T]) iter.Seq[T] {
	return func(seq iter.Seq[T]) iter.Seq[T] {
		return gofrs.Filter(seq, pred)
	}
}
