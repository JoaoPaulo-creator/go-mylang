package token

type TokenType string

const (
	LEFT_PAREN    = "("
	RIGHT_PAREN   = ")"
	LEFT_BRACE    = "{"
	RIGHT_BRACE   = "}"
	RIGHT_BRACKET = "]"
	LEFT_BRACKET  = "["
	COMMA         = ","
	DOT           = "."
	MINUS         = "-"
	PLUS          = "+"
	SEMICOLON     = ";"
	SLASH         = "/"
	STAR          = "*"

	// Tokens de um ou dois Caracteres
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Tokens literais
	IDENTIFIER
	STRING
	NUMBER

	// Palavras-chave
	AND
	CLASS
	ELSE
	MY_FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	MY_TRUE
	VAR
	WHILE

	// Token para final de declaração
	EOF
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}
