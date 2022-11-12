package tokens

type TokenPrecedence uint16

const (
	TOKEN_PRECEDENCE_OP_LPAREN = 3
	TOKEN_PRECEDENCE_OP_MULT   = 2
	TOKEN_PRECEDENCE_OP_DIV    = 2
	TOKEN_PRECEDENCE_OP_ADD    = 1
	TOKEN_PRECEDENCE_OP_SUB    = 1
)

func Precedence(t Token) TokenPrecedence {
	switch t.TokenType {
	case TOKEN_TYPE_OP:
		return precedenceOp(OperatorType(t.Text))
	}

	return 0
}

func precedenceOp(text OperatorType) TokenPrecedence {
	switch text {
	case OP_TYPE_PLUS:
		return TOKEN_PRECEDENCE_OP_ADD
	case OP_TYPE_MINUS:
		return TOKEN_PRECEDENCE_OP_SUB
	case OP_TYPE_MULT:
		return TOKEN_PRECEDENCE_OP_MULT
	case OP_TYPE_DIV:
		return TOKEN_PRECEDENCE_OP_DIV
	case OP_TYPE_LPAREN:
		return TOKEN_PRECEDENCE_OP_LPAREN
	}

	return 0
}
