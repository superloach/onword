package value

import "strconv"

type String string

func (s String) String() string {
	return string(s)
}

func (s String) Int() (int, error) {
	return strconv.Atoi(string(s))
}
