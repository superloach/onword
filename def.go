package onword

import "fmt"

// Def is a function that manipulates an OnWord instance.
type Def func(*OnWord) error

// Hint retrieves the Def's Hint, using the convention of passing a nil OnWord.
func (d Def) Hint() *Hint {
	if h, ok := d(nil).(*Hint); ok {
		return h
	}

	return nil
}

var (
	// DefAdd is a Def that adds the top 2 elements on the stack.
	DefAdd = func(o *OnWord) error {
		if o == nil {
			return NewHint(`n1 n2 -- n1 + n2`)
		}

		n2, err := o.StackPop()
		if err != nil {
			return fmt.Errorf("add pop n2: %w", err)
		}

		n1, err := o.StackPop()
		if err != nil {
			return fmt.Errorf("add pop n1: %w", err)
		}

		o.StackPush(n1 + n2)
		return nil
	}

	// DefSub is a Def that subtracts the top 2 elements on the stack.
	DefSub = func(o *OnWord) error {
		if o == nil {
			return NewHint(`n1 n2 -- n1 - n2`)
		}

		n2, err := o.StackPop()
		if err != nil {
			return fmt.Errorf("sub pop n2: %w", err)
		}

		n1, err := o.StackPop()
		if err != nil {
			return fmt.Errorf("sub pop n1: %w", err)
		}

		o.StackPush(n1 - n2)
		return nil
	}

	// DefMul is a Def that multiplies the top 2 elements on the stack.
	DefMul = func(o *OnWord) error {
		if o == nil {
			return NewHint(`n1 n2 -- n1 * n2`)
		}

		n2, err := o.StackPop()
		if err != nil {
			return fmt.Errorf("mul pop n2: %w", err)
		}

		n1, err := o.StackPop()
		if err != nil {
			return fmt.Errorf("mul pop n1: %w", err)
		}

		o.StackPush(n1 * n2)
		return nil
	}

	// DefDiv is a Def that divides the top 2 elements on the stack.
	DefDiv = func(o *OnWord) error {
		if o == nil {
			return NewHint(`n1 n2 -- n1 / n2`)
		}

		n2, err := o.StackPop()
		if err != nil {
			return fmt.Errorf("div pop n2: %w", err)
		}

		n1, err := o.StackPop()
		if err != nil {
			return fmt.Errorf("div pop n1: %w", err)
		}

		o.StackPush(n1 / n2)
		return nil
	}
)
