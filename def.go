package onword

type Def interface {
	Hint() string
	Call(*Ses) error
}
