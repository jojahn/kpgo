package rational

import (
	"fmt"
	"testing"
)

func TestRational(t *testing.T) {
	ratA := New(1,2)
	ratB := New(2,4)
	if ratA != ratB {
		t.Error("Rationals are not equal. Shortened did not work")
	}

	ratC := ratA.Multiply(New(1,2))
	fmt.Println(ratC)
	if ratC.nominator != 1 || ratC.denominator != 4 {
		t.Error("Multiply did not work")
	}

	ratC = ratA.Add(ratB)
	fmt.Println(ratC)
	if ratC.nominator != 1 || ratC.denominator != 1 {
		t.Error("Add did not work")
	}
}
