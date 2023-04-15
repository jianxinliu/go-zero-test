package util

type Predict[T any] interface {
	Test(t T) bool
}

type NotZero[T int | int8 | int16 | int32 | int64] struct {
	Predict[T]
}

func (nz NotZero[T]) Test(t T) bool {
	return t != 0
}
func IfThen[T any](a T, cond Predict[T], then func(a T)) {
	if cond.Test(a) {
		then(a)
	}
}
