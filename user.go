package onword

import "fmt"

type UserDef struct {
	Name string
	Defs []Def
}

func (s *UserDef) Hint() string {
	if len(s.Defs) == 0 {
		return "???"
	}

	hint, ok := s.Defs[0].(HintDef)
	if ok {
		return string(hint)
	}

	return "???"
}

func (s *UserDef) Call(ses *Ses) error {
	for i, d := range s.Defs {
		err := d.Call(ses)
		if err != nil {
			err = fmt.Errorf("call %d: %w", i, err)
			return err
		}
	}

	return nil
}
