package parser

import (
	"fmt"
	"github.com/jojahn/kpgo/4_Parsers/lexer"
	"testing"
)

func TestParser(t *testing.T) {
	input := "A & B | !C"
	lexer, _ := lexer.NewLexer(input)
	parser := NewParser(lexer)
	fmt.Println(parser)
	res := parser.Eval(map[string]bool{"A": true, "B": false, "C": true})
	if res != false {
		t.Error("'True & False | !True' should be false")
	}
	res = parser.Eval(map[string]bool{"A": false, "B": false, "C": false})
	if res != true {
		t.Error("'False & False | !False' should be true")
	}
}

func TestEval(t *testing.T) {
	// A & B | !C
	input := "A & B | !C"
	lexer, _ := lexer.NewLexer(input)
	parser := NewParser(lexer)
	result := parser.Eval(map[string]bool{
		"A": true,
		"B": false,
		"C": false,
	})
	if result != true {
		t.Error("expected true")
	}

	result = parser.Eval(map[string]bool{
		"A": true,
		"B": false,
		"C": true,
	})
	if result != false {
		t.Error("expected false")
	}
}
