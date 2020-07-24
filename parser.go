package onword

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/regex"
)

// Lexer defines tokens for OnWord's syntax.
var Lexer = lexer.Must(regex.New(`
	Space = [\n\r\t ]+
	Int = [0-9]+
	DefOpen = [:]
	DefClose = [;]
	HintOpen = [(]
	HintClose = [)]
	Word = [^\n\r\t ]+
`))

// Parser combines the Word grammar and the Lexer.
var Parser = participle.MustBuild(
	&Word{},
	participle.Lexer(Lexer),
	participle.Elide("Space"),
)
