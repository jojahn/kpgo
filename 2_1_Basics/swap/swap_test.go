package swap

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	x, y := Swap(1, 2)
	if x != 2 || y != 1 {
		t.Error("err")
	}
}

func TestSwapPointers(t *testing.T) {
	x, y := 1, 2
	xp := &x
	SwapPointers(&x, &y)
	xp = nil
	fmt.Println(xp)
	fmt.Println(x)
	if x != 2 || y != 1 {
		t.Error("err")
	}

	SwapPointers(nil, &x)
}