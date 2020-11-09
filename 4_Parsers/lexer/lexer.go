package lexer

import "strings"

/*
	<expression> ::= <term> { <or> <term> }
	<term> ::= <factor> { <and> <factor> }
	<factor> ::= <var> | <not> <factor> | (<expression>)
	<or>  ::= '|'
	<and> ::= '&'
	<not> ::= '!'
	<var> ::= '[a-zA-Z0-9]+'
 */

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Lexer struct {
	Iterator
	tokens []string
}

func NewLexer(input string) Lexer {
	lexer := Lexer{}
	lexer = *lexer.splitTokens(input)
	return lexer
}

func (l *Lexer) splitTokens(input string) *Lexer {
	var tokens []string = strings.Split(input, " ")
	l.tokens = tokens
	return l
}

func (l Lexer) HasNext() bool {
	return l.tokens != nil && len(l.tokens) > 0
}

func (l *Lexer) Next() string {
	if l.HasNext() {
		return l.tokens[0]
	}
	return nil
}
