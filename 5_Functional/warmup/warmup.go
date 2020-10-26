package warmup

import "fmt"

func inc(value int) int {
	return value + 1
}

func transform(function func(int) int, value int) {
	fmt.Println(function(value))
}

func multiplyC(x int) func(value int) int {
	return func(value int) int {
		value *= x
		return value
	}
}
