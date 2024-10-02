package lexer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOKTokenizer(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"1 + 2", []string{"1", "+", "2"}},
		{"(1 + 2)", []string{"(", "1", "+", "2", ")"}},
		{"(1 + (2 - 3))", []string{"(", "1", "+", "(", "2", "-", "3", ")", ")"}},
		{"-(-1 + (2- 3*7))", []string{"-", "(", "-", "1", "+", "(", "2", "-", "3", "*", "7", ")", ")"}},
	}
	for i := range tests {
		res, _ := Tokenize(tests[i].input)
		assert.Equal(t, res, tests[i].expected, "must be equal")
	}
}

func TestFailTokenize(t *testing.T) {
	tests := []string{
		"#-5+4",
		"a+b+c/4",
		"(3+2",
		"((9+2 / 4)))",
	}
	for i := range tests {
		_, err := Tokenize(tests[i])
		assert.NotNil(t, err)
	}
}
