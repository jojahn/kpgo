package parser

import (
	"fmt"
	"github.com/jojahn/kpgo/4_Parsers/lexer"
)

/*

<expression> ::= <term> { <or> <term> }
<term> ::= <factor> { <and> <factor> }
<factor> ::= <var> | <not> <factor> | (<expression>)

*/

type Evaluator interface {
	Eval(vars map[string]bool) bool
}

type or struct {
	Evaluator
	left Evaluator
	right Evaluator
}

func (o or) String() string {
	return fmt.Sprintf("%s <or> %s", o.left, o.right)
}

func (o or) Eval(vars map[string]bool) bool {
	return o.left.Eval(vars) || o.right.Eval(vars)
}

type and struct {
	Evaluator
	left Evaluator
	right Evaluator
}

func (a and) String() string {
	return fmt.Sprintf("%s <and> %s", a.left, a.right)
}

func (a and) Eval(vars map[string]bool) bool {
	return a.left.Eval(vars) && a.right.Eval(vars)
}

type not struct {
	Evaluator
	variable Evaluator
}

func (n not) String() string {
	return fmt.Sprintf("<not> %s", n.variable)
}

func (n not) Eval(vars map[string]bool) bool {
	return !n.variable.Eval(vars)
}

type varNode struct {
	Evaluator
	name string
}

func (v varNode) String() string {
	return fmt.Sprintf("<var:\"%s\">", v.name)
}

func (v varNode) Eval(vars map[string]bool) bool {
	return vars[v.name]
}

type Parser struct {
	Evaluator
	root Evaluator
	lexer lexer.Lexer
}

func NewParser(lexer lexer.Lexer) Parser {
	parser := Parser{lexer: lexer}
	parser.root = parser.expression()
	return parser
}

func (p Parser) String() string {
	return fmt.Sprintf("%s", p.root)
}

// <expression> ::= <term> { <or> <term> }
func (p *Parser) expression() Evaluator {
	left := p.term()
	token, err := p.lexer.Next()
	if err == nil && token == "|" {
		right := p.term()
		return or{left: left, right: right}
	} else {
		return left
	}
}

// <term> ::= <factor> { <and> <factor> }
func (p *Parser) term() Evaluator {
	left := p.factor()
	token, err := p.lexer.Next()
	if err == nil && token == "&" {
		right := p.factor()
		return and{left: left, right: right}
	} else {
		return left
	}
}

// <factor> ::= <var> | <not> <factor> | (<expression>)
func (p *Parser) factor() Evaluator {
	token, err := p.lexer.Next()
	if err == nil && token == "!" {
		return not{variable: p.factor()}
	} else {
		return varNode{name: token}
	}
}

func (p *Parser) Eval(vars map[string]bool) bool {
	return p.root.Eval(vars)
}