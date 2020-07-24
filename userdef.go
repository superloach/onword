package onword

import (
	"fmt"
	"strings"
)

type UserDef struct {
	Name string  `parser:"DefOpen @Word"`
	Body []*Word `parser:"@@* DefClose"`
}

func (u *UserDef) String() string {
	parts := []string{":", u.Name}

	for _, word := range u.Body {
		parts = append(parts, word.String())
	}

	parts = append(parts, ";")

	return strings.Join(parts, " ")
}

func (u *UserDef) Hint() *Hint {
	if len(u.Body) > 0 {
		return u.Body[0].Hint
	}

	return nil
}

func (u *UserDef) Def() Def {
	return func(o *OnWord) error {
		if o == nil {
			return u.Hint()
		}

		for _, word := range u.Body {
			err := o.Apply(word)
			if err != nil {
				return fmt.Errorf("userdef apply %q: %w", word, err)
			}
		}

		return nil
	}
}
