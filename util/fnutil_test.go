package util

import "testing"

func TestFnUtil(t *testing.T) {
	num := 23
	IfThen(num, NotZeroImpl[int]{}, func(i int) {
		println("not zero")
	})
}
