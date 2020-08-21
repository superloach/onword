package onword

type HintDef string

func (h HintDef) Hint() string {
	return string(h)
}

func (h HintDef) Call(_ *Ses) error {
	return nil
}
