package onword

import (
	"fmt"

	"github.com/superloach/onword/value"
)

var Builtin = map[string]Def{
	"foo": &NativeDef{" -- bar", func(ses *Ses) error {
		ses.Queue.Push(value.String("bar"))
		return nil
	}},

	"sum": &NativeDef{"a b -- a+b", func(ses *Ses) error {
		b, err := ses.Queue.Pop().Int()
		if err != nil {
			return fmt.Errorf("pop int b: %w", err)
		}

		a, err := ses.Queue.Pop().Int()
		if err != nil {
			return fmt.Errorf("pop int a: %w", err)
		}

		ses.Queue.Push(value.Int(a + b))
		return nil
	}},

	"?": &NativeDef{"word -- hint", func(ses *Ses) error {
		word := ses.Queue.Pop().String()

		d, ok := ses.Onword.Defs[word]
		if !ok {
			return fmt.Errorf("word %q not found", word)
		}

		ses.Queue.Push(value.String(d.Hint()))
		return nil
	}},

	".": &NativeDef{"value -- [prints value]", func(ses *Ses) error {
		ses.Output <- ses.Queue.Pop().String()
		return nil
	}},

	":": &NativeDef{"x -- x x", func(ses *Ses) error {
		x := ses.Queue.Pop()

		ses.Queue.Push(x)
		ses.Queue.Push(x)

		return nil
	}},
}
