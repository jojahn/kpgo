package composition

import (
	"fmt"
	"testing"
)

func TestComposition(t *testing.T) {
	fmt.Printf("%v\n", compose(square, square)(2)) // --> 4*4 = 16
	fmt.Printf("%v\n", compose(compose(square, square), square)(2)) // --> 256

	if compose(square, square)(2) != 16 {
		t.Error("compose(square, square)(2) or (2x2)x(2x2) should be 4")
	}

	if compose(square, square)(0) != 0 {
		t.Error("compose(square, square)(0) or (0x0)x(0x0) should be 0")
	}
}
