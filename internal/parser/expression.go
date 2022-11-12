package parser

import (
	"awesomeProject/internal/tokens"
)

func ParseExpression(toks []tokens.Token) []tokens.Token {
	opStack, outpQueue := []tokens.Token{}, []tokens.Token{}
	for i := 0; i < len(toks); i++ {
		token := toks[i]
		if token.TokenType == tokens.TOKEN_TYPE_NUMBER {
			push(&outpQueue, token)
		} else if token.TokenType == tokens.TOKEN_TYPE_OP {
			if token.Text == tokens.OP_TYPE_LPAREN {
				push(&opStack, token)
			} else if token.Text == tokens.OP_TYPE_RPAREN {
				for peekStack(opStack).Text != tokens.OP_TYPE_LPAREN {
					op := popStack(&opStack)
					push(&outpQueue, op)
				}
				popStack(&opStack)
			} else {
				for len(opStack) > 0 && tokens.Precedence(peekStack(opStack)) >= tokens.Precedence(token) {
					op := popStack(&opStack)
					push(&outpQueue, op)
				}
				push(&opStack, token)
			}
		}
	}

	for len(opStack) > 0 {
		op := popStack(&opStack)
		push(&outpQueue, op)
	}
	return outpQueue
}

func push(collection *[]tokens.Token, tok tokens.Token) {
	*collection = append(*collection, tok)
}

func popStack(st *[]tokens.Token) tokens.Token {
	if len(*st) == 0 {
		panic("pop on empty stack")
	}
	out := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return out
}

func peekStack(st []tokens.Token) tokens.Token {
	return st[len(st)-1]
}

func popQueue(q *[]tokens.Token) tokens.Token {
	if len(*q) == 0 {
		panic("pop on empty stack")
	}
	out := (*q)[0]
	*q = (*q)[1:]
	return out
}

func peekQueue(q []tokens.Token) tokens.Token {
	return q[0]
}
