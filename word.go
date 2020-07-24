package onword

import "fmt"

// Word is a basic parsing element containing a Word or an Int.
type Word struct {
	Hint    *Hint    `parser:"@@"`
	Word    *string  `parser:"| @Word"`
	Int     *int     `parser:"| @Int"`
	UserDef *UserDef `parser:"| @@"`
}

// String returns the Word's string representation.
func (w *Word) String() string {
	if w == nil {
		return "<nil>"
	}

	if w.Word != nil {
		return *w.Word
	}

	if w.Int != nil {
		return fmt.Sprintf("%d", *w.Int)
	}

	if w.Hint != nil {
		return w.Hint.String()
	}

	if w.UserDef != nil {
		return w.UserDef.String()
	}

	return "<nil>"
}
