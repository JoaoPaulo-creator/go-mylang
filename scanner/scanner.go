package scanner

import "my-lang/token"

type Scanner struct {
	start   int
	current int
	line    int
	source  string
	tokens  []token.Token
}

var keywords = map[string]token.TokenType{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"false":  token.MY_FALSE,
	"for":    token.FOR,
	"fun":    token.FUN,
	"if":     token.IF,
	"nil":    token.NIL,
	"or":     token.OR,
	"print":  token.PRINT,
	"return": token.RETURN,
	"super":  token.SUPER,
	"this":   token.THIS,
	"true":   token.MY_TRUE,
	"var":    token.VAR,
	"while":  token.WHILE,
}
