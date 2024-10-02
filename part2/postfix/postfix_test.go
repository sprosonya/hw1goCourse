package postfix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOKCalculatePostfix(t *testing.T) {
	tests := []struct {
		input    []string
		expected float64
	}{
		{[]string{"1", "2", "+"}, 3},
		{[]string{"3", "4", "*"}, 12},
		{[]string{"5", "3", "-"}, 2},
		{[]string{"8", "2", "/"}, 4},
		{[]string{"5", "~"}, -5},
		{[]string{"5", "1", "2", "~", "-", "/"}, 5.0 / 3},
		{[]string{"11", "~", "1", "20", "*", "2", "/", "-", "11", "2", "/", "3", "*", "-", "~"}, 37.5},
	}
	for i := range tests {
		res, _ := CalculatePostfix(tests[i].input)
		assert.Equal(t, res, tests[i].expected, "must be equal")
	}
}
func TestFailCalculatePostfix(t *testing.T) {
	tests := [][]string{
		[]string{"8", "0", "/"},
		[]string{"1", "+"},
		[]string{"1", "a", "+"},
		[]string{"1", "2", "+", "3"},
	}
	for i := range tests {
		_, err := CalculatePostfix(tests[i])
		assert.NotNil(t, err)
	}
}
