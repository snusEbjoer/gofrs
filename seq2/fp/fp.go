package fp

import (
	"iter"

	"github.com/snusEbjoer/gofrs/seq2"
)

func Map[T1, T2, R1, R2 any](f func(T1, T2) (R1, R2)) func(iter.Seq2[T1, T2]) iter.Seq2[R1, R2] {
	return func(s iter.Seq2[T1, T2]) iter.Seq2[R1, R2] {
		return seq2.Map(s, f)
	}
}

func MapToSeq1[T1, T2, R any](f func(T1, T2) R) func(iter.Seq2[T1, T2]) iter.Seq[R] {
	return func(s iter.Seq2[T1, T2]) iter.Seq[R] {
		return seq2.MapToSeq1(s, f)
	}
}

func FilterMap[T1, T2, R1, R2 any](f func(T1, T2) (R1, R2, bool)) func(iter.Seq2[T1, T2]) iter.Seq2[R1, R2] {
	return func(s iter.Seq2[T1, T2]) iter.Seq2[R1, R2] {
		return seq2.FilterMap(s, f)
	}
}

func FilterMapToSeq1[T1, T2, R any](f func(T1, T2) (R, bool)) func(iter.Seq2[T1, T2]) iter.Seq[R] {
	return func(s iter.Seq2[T1, T2]) iter.Seq[R] {
		return seq2.FilterMapToSeq1(s, f)
	}
}
