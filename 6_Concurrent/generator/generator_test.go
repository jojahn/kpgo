package generator

import (
	"testing"
	"fmt"
)

func TestGenerator(t *testing.T) {
	fibChan := fib()
	for n := 1; n <= 10; n++ {
		fmt.Printf("n=%d, fib=%d\n", n, <-fibChan)
	}
}