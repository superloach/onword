package value

type Nil struct{}

func (n Nil) String() string {
	return "<nil>"
}

func (n Nil) Int() (int, error) {
	return 0, nil
}
