package scanner

import (
	"errors"
)

type Scanner struct {
	source string
	tokens []Token

	start   int
	current int
	line    int
}

func (s *Scanner) Init(source string) {
	s.source = source
}

func (s *Scanner) ScanTokens() {
	for {
		if s.isAtEnd() {
			break
		}

		s.start = s.current
		s.scanToken()
	}
}

func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() error {
	c := s.advance()

	switch c {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '\\':
		s.addToken(SLASH)
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			// TODO check if keyword or identifier
		}

		return errors.New("Invalid token " + string(c))
	}

	return nil
}

func (s *Scanner) number() {
	// TODO
}

func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current += 1
	return c
}

func (s Scanner) lookahead() byte {
	if s.isAtEnd() {
		return '\x00'
	}

	return s.source[s.current]
}

func (s *Scanner) addToken(token Token) {
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) consume() {
	s.current += 1
}

func (s *Scanner) match(c byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] == c {
		s.current += 1
		return true
	}

	return false
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isAlphanumeric(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
