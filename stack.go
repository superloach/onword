package onword

import "fmt"

// StackPush places an int at the top of the OnWord's stack.
func (o *OnWord) StackPush(n int) {
	o.Stack = append(o.Stack, n)
}

// StackPop removes an int from the top of the OnWord's stack.
func (o *OnWord) StackPop() (int, error) {
	if len(o.Stack) == 0 {
		return 0, fmt.Errorf("stack pop: %w", ErrOutOfBounds)
	}

	n := (o.Stack)[len(o.Stack)-1]
	o.Stack = (o.Stack)[:len(o.Stack)-1]

	return n, nil
}
