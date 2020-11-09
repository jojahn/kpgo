package lexer

import (
	"fmt"
	"regexp"
	"strings"
)

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

func NewLexer(input string) (Lexer, error) {
	tokens, err := splitTokens(input)
	if err != nil {
		return Lexer{}, err
	}
	return Lexer{tokens: tokens}, nil
}

func (l Lexer) String() string {
	return fmt.Sprintf("Lexer {\n tokens %s \n}", l.tokens)
}

func splitTokens(input string) ([]string, error) {
	tokens := make([]string, 0)
	for _, char := range input {
		switch char {
		case '&', '|', '!':
			tokens = append(tokens, fmt.Sprintf("%c", char))
		case 0x20:
			continue
		default:
			isLetter := char >= 'a' && char <= 'z'
			isCapitalLetter := char >= 'A' && char <= 'Z'
			isNumber := char >= '0' && char <= '9'
			if isCapitalLetter || isLetter || isNumber {
				tokens = append(tokens, fmt.Sprintf("%c", char))
			} else {
				return make([]string, 0), fmt.Errorf("InvalidToken")
			}
		}
	}
	return tokens, nil
}

func splitTokensSimple(input string) ([]string, error) {
	var tokens []string = make([]string, 0)
	for _, token := range strings.Split(input, " ") {
		if token[0] == '!' {
			token = token[1:]
			tokens = append(tokens, "!")
		}
		if isValid(token) == false {
			return make([]string, 0), fmt.Errorf("InvalidToken")
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func isValid(token string) bool {
	allowedTokens := []string{
		"^\\|$",
		"^\\&$",
		"^!$",
		"^[a-zA-Z0-9]+$",
	}

	for _, allowedToken := range allowedTokens {
		valid, err := regexp.Match(allowedToken, []byte(token))
		if err != nil {
			return false
		}

		if valid == true {
			return true
		}
	}
	return false
}

func (l Lexer) HasNext() bool {
	return l.tokens != nil && len(l.tokens) > 0
}

func (l *Lexer) Next() (string,  error) {
	if l.HasNext() {
		token := l.tokens[0]
		l.tokens = l.tokens[1:]
		return token, nil
	}
	return "", fmt.Errorf("ArrayOutOfBounds")
}
