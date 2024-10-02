package stack

type Stack[T any] struct {
	Items []T
}

func NewEmptyStack[T any]() *Stack[T] {
	return &Stack[T]{
		Items: nil,
	}
}

func NewStack[T any](Items []T) *Stack[T] {
	return &Stack[T]{
		Items: Items,
	}
}

func (stack *Stack[T]) Push(item T) {
	stack.Items = append(stack.Items, item)
}

func (stack *Stack[T]) Pop() (lastItem T) {
	if len(stack.Items) == 0 {
		return
	}
	lastItem = stack.Items[len(stack.Items)-1]
	stack.Items = stack.Items[:len(stack.Items)-1]
	return lastItem
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.Items) == 0
}

func (stack *Stack[T]) Top() (lastItem T) {
	if len(stack.Items) == 0 {
		return
	}
	return stack.Items[len(stack.Items)-1]
}

func (stack *Stack[T]) Size() int {
	return len(stack.Items)
}

func (stack *Stack[T]) Clear() {
	stack.Items = nil
}
