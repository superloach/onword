package onword

import "github.com/superloach/onword/value"

type Queue []value.Value

func (q *Queue) Push(v value.Value) {
	*q = append(*q, v)
}

func (q *Queue) Pop() value.Value {
	if len(*q) == 0 {
		return value.Nil{}
	}

	v := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]

	return v
}
