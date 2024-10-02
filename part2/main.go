package main

import (
	"bufio"
	"calc/dijkstra"
	"calc/lexer"
	"calc/postfix"
	"fmt"
	"os"
)

func main() {
	var input string
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input = s.Text()

	tokens, err := lexer.Tokenize(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	postfixExpression := dijkstra.InfixToPostfix(tokens)

	num, err := postfix.CalculatePostfix(postfixExpression)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Результат:", num)
}
