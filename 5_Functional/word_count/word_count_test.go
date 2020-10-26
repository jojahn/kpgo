package word_count

import (
	"fmt"
	"github.com/jojahn/kpgo/5_Functional/map_filter_reduce"
	"testing"
)

func TestWordCount(t *testing.T) {
	strings := []map_filter_reduce.Any {"a", "a", "b", "c", "A", "a", "b", "abc"}
	pairs := countWords(strings)
	fmt.Println(pairs)
	if pairs[0].count != 3 {
		t.Error("a should have 3 occurrences")
	}
	if pairs[0].word != "a" {
		t.Error("a should be the first word")
	}
}
