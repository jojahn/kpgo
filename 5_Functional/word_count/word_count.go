package word_count

import (
	"fmt"
	"github.com/jojahn/kpgo/5_Functional/map_filter_reduce"
)

type Pair struct {
	word string
	count int
}

func (p Pair) String() string {
	return fmt.Sprintf("(%s,%d)", p.word, p.count)
}

func countWords(strings []map_filter_reduce.Any) []Pair {
	return map_filter_reduce.ToStream(strings).
		Map(toPair).
		Reduce(sumPairs).([]Pair)
}

func sumPairs(acc, value map_filter_reduce.Any) map_filter_reduce.Any {
	if acc == nil {
		acc = []Pair{}
	}
	for i, item := range acc.([]Pair) {
		if item.word == value.(Pair).word {
			acc.([]Pair)[i].count++
			return acc
		}
	}
	return append(acc.([]Pair), value.(Pair))
}

func toPair(s map_filter_reduce.Any) map_filter_reduce.Any {
	return Pair{s.(string), 1}
}