package postfix

import (
	"calc/stack"
	"fmt"
	"strconv"
)

func CalculatePostfix(items []string) (float64, error) {
	var stack *stack.Stack[float64] = stack.NewEmptyStack[float64]()
	var num float64
	var err error
	// Errors
	DivisionByZero := "Error: division by zero"
	NotEnoughOperands := "Error: not enough operands for operand %s : %v\n"
	NotAValidNumber := "Error: %s is not a valid number\n"
	NotOneNumber := "Error: result consists not of one number : %v\n"

	for _, item := range items {
		switch item {
		// Binary operators
		case "+", "-", "*", "/":
			if stack.Size() < 2 {
				return 0, fmt.Errorf(NotEnoughOperands, item, stack.Items)
			}
			switch item {
			case "+":
				stack.Push(stack.Pop() + stack.Pop())
			case "-":
				stack.Push(-stack.Pop() + stack.Pop())
			case "*":
				stack.Push(stack.Pop() * stack.Pop())
			case "/":
				var divider float64 = stack.Pop()
				if divider == 0 {
					return 0, fmt.Errorf(DivisionByZero)
				}
				stack.Push(stack.Pop() / divider)
			}
		// Unary operators
		case "~":
			if stack.Size() < 1 {
				return 0, fmt.Errorf(NotEnoughOperands, item, stack.Items)
			}
			stack.Push(-stack.Pop())
		// Numbers
		default:
			num, err = strconv.ParseFloat(item, 64)
			if err != nil {
				return 0, fmt.Errorf(NotAValidNumber, item)
			}
			stack.Push(num)
		}
	}
	if stack.Size() != 1 {
		return 0, fmt.Errorf(NotOneNumber, stack.Items)
	}
	return stack.Pop(), nil
}
