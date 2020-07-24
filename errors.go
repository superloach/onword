package onword

import "errors"

var (
	// ErrNilWord occurs when a Word has no String or Number part.
	ErrNilWord = errors.New("nil word")

	// ErrUnknownWord occurs when a Word's String part is not known.
	ErrUnknownWord = errors.New("unknown word")

	// ErrOutOfBounds indicates that an index outside of bounds was used.
	ErrOutOfBounds = errors.New("out of bounds")
)
