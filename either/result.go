package either

type Result[T any] = Either[T, error]

func NewResult[T any](v T, err error) Result[T] {
	if err != nil {
		return Result[T]{
			right:  err,
			isLeft: false,
		}
	}
	return Result[T]{
		left:   v,
		isLeft: true,
	}
}
