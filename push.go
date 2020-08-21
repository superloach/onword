package onword

import (
	"fmt"

	"github.com/superloach/onword/value"
)

type PushDef struct {
	value.Value
}

func (p PushDef) Hint() string {
	return fmt.Sprintf("( -- %s)", p)
}

func (p PushDef) Call(ses *Ses) error {
	ses.Queue.Push(p.Value)
	return nil
}
