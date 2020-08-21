package onword

type NativeDef struct {
	H string
	C func(*Ses) error
}

func (n NativeDef) Hint() string {
	return n.H
}

func (n NativeDef) Call(ses *Ses) error {
	return n.C(ses)
}
