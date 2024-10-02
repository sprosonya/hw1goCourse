package lexer

import (
	"calc/stack"
	"calc/tokens"
	"fmt"
)

type Lexer struct {
	input  string
	pos    int
	ch     byte
	length int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input, pos: -1, length: len(input)}
	l.NextChar()
	return l
}

func (l *Lexer) NextChar() {
	l.pos++
	if l.pos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.pos]
	}
}

func (l *Lexer) Next() bool {
	return l.pos < l.length
}

func (l *Lexer) ProcessNextChar(masOfTokens *[]string, brackets *stack.Stack[byte]) error {
	for l.ch == ' ' {
		l.NextChar()
	}

	switch {
	case tokens.IsOperator(string(l.ch)):
		*masOfTokens = append(*masOfTokens, string(l.ch))
		l.NextChar()
	case l.ch == '(':
		*masOfTokens = append(*masOfTokens, string(l.ch))
		brackets.Push(l.ch)
		l.NextChar()
	case l.ch == ')':
		*masOfTokens = append(*masOfTokens, string(l.ch))
		if brackets.IsEmpty() {
			return fmt.Errorf("error with brackets")
		}
		_ = brackets.Pop()
		l.NextChar()
	case tokens.IsDigit(l.ch):
		start := l.pos
		for tokens.IsDigit(l.ch) {
			l.NextChar()
		}
		*masOfTokens = append(*masOfTokens, l.input[start:l.pos])
	default:
		return fmt.Errorf("unknown symbol")
	}
	return nil
}

func Tokenize(input string) ([]string, error) {
	lexer := NewLexer(input)
	tokens := make([]string, 0)
	brackets := stack.NewEmptyStack[byte]()

	for lexer.Next() {
		if err := lexer.ProcessNextChar(&tokens, brackets); err != nil {
			return nil, err
		}
	}
	if !brackets.IsEmpty() {
		return nil, fmt.Errorf("error with brackets")
	}
	return tokens, nil
}
