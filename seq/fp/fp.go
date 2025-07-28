package fp

import (
	"iter"

	"github.com/snusEbjoer/gofrs/helpers"
	"github.com/snusEbjoer/gofrs/seq"
)

func Map[T, R any](f func(T) R) func(iter.Seq[T]) iter.Seq[R] {
	return func(s iter.Seq[T]) iter.Seq[R] {
		return seq.Map(s, f)
	}
}

func MapToSeq2[T, R1, R2 any](f func(T) (R1, R2)) func(iter.Seq[T]) iter.Seq2[R1, R2] {
	return func(s iter.Seq[T]) iter.Seq2[R1, R2] {
		return seq.MapToSeq2(s, f)
	}
}

func FilterMapToSeq2[T, R1, R2 any](f func(T) (R1, R2, bool)) func(iter.Seq[T]) iter.Seq2[R1, R2] {
	return func(s iter.Seq[T]) iter.Seq2[R1, R2] {
		return seq.FilterMapToSeq2(s, f)
	}
}

func FilterMap[T, R any](f func(T) (R, bool)) func(iter.Seq[T]) iter.Seq[R] {
	return func(s iter.Seq[T]) iter.Seq[R] {
		return seq.FilterMap(s, f)
	}
}

func DistinctMap[T any, R comparable](f func(T) R) func(iter.Seq[T]) iter.Seq[R] {
	return helpers.Comp2(Map(f), seq.Distinct)
}
