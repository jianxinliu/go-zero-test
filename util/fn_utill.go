package util

type Number interface {
	int | int8 | int16 | int32 | int64
}

type Predict[T any] interface {
	Test(t T) bool
}

type NotZero[T Number] interface {
	Predict[T]
}

type NotZeroImpl[T Number] struct {
}

func (nz NotZeroImpl[T]) Test(t T) bool {
	return t != 0
}

func IfThen[T any](a T, cond Predict[T], then func(a T)) {
	if cond.Test(a) {
		then(a)
	}
}
