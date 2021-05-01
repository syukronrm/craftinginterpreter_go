package scanner

import (
	"testing"
)

func TestScanner(t *testing.T) {
	s := new(Scanner)
	s.Init("(")
	s.ScanTokens()
	if len(s.tokens) != 1 {
		t.Fatalf(`Should be 1, got %d`, len(s.tokens))
	}

	if s.tokens[0] != LEFT_PAREN {
		t.Fatalf(`Should be %d, got %d`, LEFT_PAREN, len(s.tokens))
	}
}
