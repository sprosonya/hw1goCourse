package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"calc/lexer"
	"calc/dijkstra"
	"calc/postfix"
)

func TestMain(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"1 + (2/2 +(3 *9 - 2) - 2)", 25},
		{"(7 + (2 * (3 + 4))) / 2", 10.5},
		{"(10 - (3 + (2 * 4))) * 2", -2},
		{"((6 + 4) / 2) + ((3 * 5) - (8 / 4))", 18},
		{"-(-11-(1*20/2)-11/2*3)", 37.5},
		{"-1 * -(4+2)", 6},
	}
	for i := range tests {
		tokens, _ := lexer.Tokenize(tests[i].input)
		postfixExpr := dijkstra.InfixToPostfix(tokens)
		res, _ := postfix.CalculatePostfix(postfixExpr)
		assert.Equal(t, res, tests[i].expected, "must be equal")
	}
}