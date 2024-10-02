package dijkstra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDijkstra(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"-", "(", "-", "11", "-", "(", "1", "*", "20", "/", "2", ")", "-", "11", "/", "2", "*", "3", ")"}, []string{"11", "~", "1", "20", "*", "2", "/", "-", "11", "2", "/", "3", "*", "-", "~"}},
		{[]string{"3", "+", "(", "4", "*", "9", "-", "2", ")", "/", "4"}, []string{"3", "4", "9", "*", "2", "-", "4", "/", "+"}},
		{[]string{"3", "*", "-", "(", "4", "+", "5", ")"}, []string{"3", "4", "5", "+", "~", "*"}},
	}
	for i := range tests {
		res := InfixToPostfix(tests[i].input)
		assert.Equal(t, res, tests[i].expected, "must be equal")
	}
}
