// Package onword contains parsing and runtime portions for the OnWord language.
package onword

import "fmt"

// OnWord contains the state of an OnWord interpreter.
type OnWord struct {
	Words  map[string]Def
	Stack  []int
	Memory []int
	Mode   Mode
}

// NewOnWord creates an OnWord instance with sane defaults.
func NewOnWord() *OnWord {
	return &OnWord{
		Words: map[string]Def{
			"+": DefAdd,
			"-": DefSub,
			"*": DefMul,
			"/": DefDiv,
		},
		Stack:  []int{},
		Memory: []int{},
	}
}

// Apply executes a Word on the OnWord instance.
func (o *OnWord) Apply(w *Word) error {
	if w.Hint != nil {
		return nil
	}

	if w.Word != nil {
		def, ok := o.Words[*w.Word]
		if !ok {
			return fmt.Errorf("find word %q: %w", w, ErrUnknownWord)
		}

		return def(o)
	}

	if w.Int != nil {
		o.StackPush(*w.Int)

		return nil
	}

	if w.UserDef != nil {
		o.Words[w.UserDef.Name] = w.UserDef.Def()

		return nil
	}

	return fmt.Errorf("apply word %q: %w", w, ErrNilWord)
}
