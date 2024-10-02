package dijkstra

import (
	"calc/stack"
	"calc/tokens"
)

func InfixToPostfix(masOfTokens []string) []string {
	output := stack.NewEmptyStack[string]()
	stack := stack.NewEmptyStack[string]()

	for i := 0; i < len(masOfTokens); i++ {
		token := masOfTokens[i]
		switch {
		case tokens.IsNumber(token):
			output.Push(token)
		case token == "(":
			stack.Push(token)
		case token == ")":
			for !stack.IsEmpty() && stack.Top() != "(" {
				temp := stack.Pop()
				output.Push(temp)
			}
			_ = stack.Pop()
		default: // operator
			if token == "-" && (i == 0 || tokens.IsOperator(masOfTokens[i-1]) || masOfTokens[i-1] == "(") {
				token = "~" //symbol of unary minus
			}
			for !stack.IsEmpty() && (tokens.GetPriority(stack.Top()) >= tokens.GetPriority(token)) {
				tmp := stack.Pop()
				output.Push(tmp)
			}
			stack.Push(token)
		}
	}
	for !stack.IsEmpty() {
		tmp := stack.Pop()
		output.Push(tmp)
	}
	return output.Items
}
