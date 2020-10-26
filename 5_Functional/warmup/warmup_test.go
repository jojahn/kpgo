package warmup

import (
	"fmt"
	"testing"
)

func TestWarmup(t *testing.T) {
	inc2 := func(value int) int {
		value++
		return value
	}
	fmt.Println(inc2(1))
	transform(inc, 1)
	transform(func(value int) int {
		return value + 1
	}, 2)
	fmt.Println(multiplyC(5)(5))
}
