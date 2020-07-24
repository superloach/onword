package onword

import (
	"fmt"
	"strings"
)

type Hint struct {
	Conts []string `parser:"HintOpen @(Space|Word|Int)* HintClose"`
}

func NewHint(s string) Hint {
	return Hint{
		Conts: strings.Split(s, " "),
	}
}

func (h Hint) String() string {
	return fmt.Sprintf("( %s )", strings.Join(h.Conts, " "))
}

func (h Hint) Error() string {
	return h.String()
}
