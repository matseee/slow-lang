package ast

import (
	"slow-lang/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&SetStatement{
				Token: token.Token{Type: token.SET, Literal: "set"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "set myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got %q", program.String())
	}
}
