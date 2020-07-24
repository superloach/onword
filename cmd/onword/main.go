package main

import (
	"fmt"

	"github.com/superloach/onword"
)

func main() {
	words := make(chan *onword.Word)

	go func(code string) {
		fmt.Printf("code %q\n", code)

		err := onword.Parser.ParseString(code, words)
		if err != nil {
			panic(fmt.Errorf("parse %#q: %w", code, err))
		}
	}(`
		: foo
			( n -- n + 10 )
			10 +
		;
		5 foo
	`)

	onw := onword.NewOnWord()

	for word := range words {
		fmt.Printf("word %q\n", word)

		err := onw.Apply(word)
		if err != nil {
			panic(fmt.Errorf("apply %q: %w", word, err))
		}

		fmt.Printf("stack %v\n", onw.Stack)
	}
}
