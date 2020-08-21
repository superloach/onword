package onword

type Onword struct {
	Defs map[string]Def
	Sess map[uint]*Ses
}

func NewOnword() *Onword {
	return &Onword{
		Defs: Builtin,
		Sess: map[uint]*Ses{},
	}
}

func (o *Onword) nextID() uint {
	i := uint(0)

	for {
		_, ok := o.Sess[i]
		if !ok {
			return i
		}

		i++
	}
}

func (o *Onword) Open() *Ses {
	id := o.nextID()
	ses := &Ses{
		Onword: o,

		ID:     id,
		Queue:  &Queue{},
		Input:  make(chan string),
		Output: make(chan string),
	}

	o.Sess[id] = ses

	return ses
}
