package map_filter_reduce

import (
	"fmt"
	"testing"
)

func TestMapFilterReduce(t *testing.T) {
	stringSlice := []Any{"a", "b", "c", "1", "D"}
	// Map/Reduce
	result := ToStream(stringSlice).
		Map(toUpperCase).
		Filter(notDigit).
		Reduce(concat).(string)

	fmt.Println(result)

	if result != "A,B,C,D" {
		t.Error(fmt.Sprintf("Result should be 'A,B,C,D' but is: %v", result))
	}
}
