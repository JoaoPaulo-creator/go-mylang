package scanner

import (
	"log"
	"my-lang/token"
	"strconv"
)

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

func (s *Scanner) IsAtEnd() bool {
	return s.current > len(s.source)
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		start:   0,
		current: 0,
		line:    1,
		tokens:  []token.Token{},
	}
}

func (s *Scanner) ScanToken() []token.Token {
	for !s.IsAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, token.Token{
		Type:    token.EOF,
		Lexeme:  "",
		Literal: nil,
		Line:    s.line,
	})
	return s.tokens
}

func (s *Scanner) isAlpha(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_'
}

func (s *Scanner) isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) isAlphaNumeric(c rune) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *Scanner) match(expected rune) bool {
	if s.IsAtEnd() || s.source[s.current] != byte(expected) {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) advance() rune {
	r := rune(s.source[s.current])
	s.current++
	return r
}

func (s *Scanner) addTokenSimple(token token.TokenType) {
	s.addToken(token, nil)
}

func (s *Scanner) addToken(tokenType token.TokenType, literal any) {
	text := s.source[s.start : s.current-s.start]
	s.tokens = append(s.tokens, token.Token{
		Type:    tokenType,
		Lexeme:  text,
		Literal: literal,
		Line:    s.line,
	})
}

func (s *Scanner) peek() rune {
	if s.IsAtEnd() {
		return '0'
	}

	return rune(s.source[s.current])
}

func (s *Scanner) peekNext() rune {
	if s.current+1 > len(s.source) {
		return '0'
	}

	return rune(s.source[s.current+1])
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.IsAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.IsAtEnd() {
		log.Printf("unterminated string %d", s.line)
		return
	}

	s.advance()

	// substring
	value := s.source[s.start+1 : s.current-s.start-2]
	s.addToken(token.STRING, value)
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()
		for s.isDigit(s.peek()) {
			s.advance()
		}
	}

	text := s.source[s.start : s.current-s.start]
	number, _ := strconv.ParseFloat(text, 64)
	s.addToken(token.NUMBER, number)
}

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	tokenType := token.TokenType("IDENTIFIER")

	if tt, ok := keywords[text]; ok {
		tokenType = tt
	}

	s.addToken(tokenType, nil)
}

func (s *Scanner) skipWhitespace() {
	for s.current == ' ' || s.current == '\t' || s.current == '\n' || s.current == '\r' {
		s.advance()
	}
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addTokenSimple(token.LEFT_PAREN)
	case ')':
		s.addTokenSimple(token.RIGHT_PAREN)
	case '{':
		s.addTokenSimple(token.LEFT_BRACE)
	case '}':
		s.addTokenSimple(token.RIGHT_BRACE)
	case ',':
		s.addTokenSimple(token.COMMA)
	case '.':
		s.addTokenSimple(token.DOT)
	case '-':
		s.addTokenSimple(token.MINUS)
	case '+':
		s.addTokenSimple(token.PLUS)
	case ';':
		s.addTokenSimple(token.SEMICOLON)
	case '*':
		s.addTokenSimple(token.STAR)
	case '[':
		s.addTokenSimple(token.LEFT_BRACKET)
	case ']':
		s.addTokenSimple(token.RIGHT_BRACKET)
	case '!':
		if s.match('=') {
			s.addTokenSimple(token.BANG_EQUAL)
		} else {
			s.addTokenSimple(token.BANG)
		}
	case '=':
		if s.match('=') {
			s.addTokenSimple(token.EQUAL_EQUAL)
		} else {
			s.addTokenSimple(token.EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addTokenSimple(token.LESS_EQUAL)
		} else {
			s.addTokenSimple(token.LESS)
		}
	case '>':
		if s.match('=') {
			s.addTokenSimple(token.GREATER_EQUAL)
		} else {
			s.addTokenSimple(token.GREATER)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.IsAtEnd() {
				s.advance()
			}
		} else {
			s.addTokenSimple(token.SLASH)
		}
	case ' ', '\r', '\t':
		// Ignore whitespace
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		if s.isDigit(c) {
			s.number()
		} else if s.isAlpha(c) {
			s.identifier()
		} else {
			// Assuming Debug is a logging mechanism; replace with your error handling
			log.Printf("Line %d: Unexpected character %c", s.line, c)
		}
	}
}
