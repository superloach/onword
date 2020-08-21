package value

import "strconv"

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func (i Int) Int() (int, error) {
	return int(i), nil
}
