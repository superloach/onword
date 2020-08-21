package value

type Value interface {
	String() string
	Int() (int, error)
}
