package types

type Result[T any] struct {
	Value T
	Err   error
}

func Fail[T any](err error) Result[T] {
	return Result[T]{
		Err: err,
	}
}

func Success[T any](value T) Result[T] {
	return Result[T]{
		Value: value,
	}
}
