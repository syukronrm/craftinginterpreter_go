package scanner

type Token int

const (
	ILLEGAL     Token = iota
	LEFT_PAREN        // (
	RIGHT_PAREN       // )
	LEFT_BRACE        // {
	RIGHT_BRACE       // }
	COMMA             // ,
	DOT               // .
	MINUS             // -
	PLUS              // +
	SEMICOLON         // ;
	SLASH             // \
	STAR              // *

	BANG          // !
	BANG_EQUAL    // !=
	EQUAL         // =
	EQUAL_EQUAL   // ==
	GREATER       // >
	GREATER_EQUAL // >=
	LESS          // <
	LESS_EQUAL    // <=

	IDENTIFIER
	STRING
	NUMBER

	AND    // and
	CLASS  // class
	ELSE   // else
	FUN    // fun
	FOR    // for
	IF     // if
	NIL    // nil
	OR     // or
	PRINT  // print
	RETURN // return
	SUPER  // super
	THIS   // this
	TRUE   // true
	VAR    // var
	WHILE  // while

	EOF
)

func (t Token) ToString() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case LEFT_PAREN:
		return "("
	case RIGHT_PAREN:
		return ")"
	case LEFT_BRACE:
		return "{"
	case RIGHT_BRACE:
		return "}"
	case COMMA:
		return ","
	// TODO
	default:
		return "ILLEGAL"
	}
}
