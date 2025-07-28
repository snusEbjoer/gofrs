package option

import (
	E "github.com/snusEbjoer/gofrs/either"
	H "github.com/snusEbjoer/gofrs/helpers"
)

type Option[T any] struct {
	value  T
	isSome bool
}

func IsSome[T any](o Option[T]) bool {
	return o.isSome
}

func IsNone[T any](o Option[T]) bool {
	return !o.isSome
}

func Some[T any](v T) Option[T] {
	return Option[T]{value: v, isSome: true}
}

func None[T any]() Option[T] {
	return Option[T]{isSome: false}
}

func GetLeft[L, R any](either E.Either[L, R]) Option[L] {
	return E.Match(
		Some,
		H.Const1[R, Option[L]](None),
	)(either)
}

func GetRight[L, R any](either E.Either[L, R]) Option[R] {
	return E.Match(
		H.Const1[L, Option[R]](None),
		Some,
	)(either)
}

func FromPredicate[T any](value T, pred H.Predicate[T]) Option[T] {
	return Option[T]{
		value:  value,
		isSome: pred(value),
	}
}

func FromNullable[T any](t *T) Option[*T] {
	return FromPredicate(t, H.NotNil)
}

func Match[T, R any](onSome H.F1[T, R], onNone H.Lazy[R]) func(Option[T]) R {
	return func(o Option[T]) R {
		if o.isSome {
			return onSome(o.value)
		}
		return onNone()
	}
}

func Filter[T any](pred H.Predicate[T]) H.F1[Option[T], Option[T]] {
	return func(o Option[T]) Option[T] {
		return Option[T]{
			value:  o.value,
			isSome: pred(o.value),
		}
	}
}

func GetOrElse[T any](onNone H.Lazy[T]) func(Option[T]) T {
	return Match(H.Identity, onNone)
}

func OrElse[T any](onNone H.Lazy[Option[T]]) H.F1[Option[T], Option[T]] {
	return Match(Some, onNone)
}

func Compact[T any](o Option[Option[T]]) Option[T] {
	return Match(
		H.Identity,
		None[T],
	)(o)
}

func Map[T, R any](mapFn H.F1[T, R]) H.F1[Option[T], Option[R]] {
	return Match(
		H.Comp2(mapFn, Some),
		None[R],
	)
}

func FilterMap[T, R any](mapFn H.F1[T, Option[R]]) H.F1[Option[T], Option[R]] {
	return H.Comp2(Map(mapFn), Compact)
}
