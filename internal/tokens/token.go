package tokens

import "fmt"

type TokenType uint8

// 1 + 2 * 3 / 44 - (10 + 15)
const (
	TOKEN_TYPE_EOF = iota
	TOKEN_TYPE_NUMBER
	TOKEN_TYPE_OP
	TOKEN_TYPE_START_EXPR
	TOKEN_TYPE_END_EXPR
)

type OperatorType string

const (
	OP_TYPE_PLUS   = "+"
	OP_TYPE_MINUS  = "-"
	OP_TYPE_MULT   = "*"
	OP_TYPE_DIV    = "/"
	OP_TYPE_LPAREN = "("
	OP_TYPE_RPAREN = ")"
)

func (tt TokenType) String() string {
	switch tt {
	case TOKEN_TYPE_EOF:
		return "EOF"
	case TOKEN_TYPE_NUMBER:
		return "NUM"
	case TOKEN_TYPE_OP:
		return "OP"
	case TOKEN_TYPE_START_EXPR:
		return "START_EXPR"
	case TOKEN_TYPE_END_EXPR:
		return "END_EXPR"
	default:
		panic(fmt.Sprintf("Forgot TokenType String() case: %d", uint8(tt)))
	}
}

type Token struct {
	TokenType TokenType
	Text      string
	Col       uint
	Line      uint
}

func (t Token) String() string {
	return fmt.Sprintf("[%s]\t(%d,%d)\t%s", t.TokenType, t.Line, t.Col, t.Text)
}
