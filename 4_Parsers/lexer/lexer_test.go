package lexer

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	lexer, err := NewLexer("A & B | !C")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("[")
	for lexer.HasNext() {
		token, err := lexer.Next()
		if err != nil {
			t.Error(err)
		} else {
			fmt.Printf("\"%s\"", token)
			if lexer.HasNext() == true {
				fmt.Printf(", ")
			}
		}
	}
	fmt.Printf("]\n")

	_, err = lexer.Next()
	if err == nil {
		t.Error("Cannot access non existing token")
	}

	lexer, err = NewLexer("% & B | !C")
	if err == nil {
		t.Error("% is an invalid token")
	}

	lexer, err = NewLexer("A & B | !!C")
	if err != nil || lexer.tokens[4] != "!" || lexer.tokens[5] != "!" || lexer.tokens[6] != "C" {
		t.Error("!!X should be parsed accordingly")
	}
}