package parse

import (
	"awesomeProject/internal/ast"
	"awesomeProject/internal/lex"
	"fmt"
)

type Error struct {
	//Parent   *Node
	Found    lex.Token
	Expected []lex.TokenType
}

func (p Error) Error() string {
	return fmt.Sprintf("Found (%s) but expected one of %s.", p.Found, p.Expected)
}

type ExprError struct {
	Found   lex.Token
	Message string
}

func (p ExprError) Error() string {
	return fmt.Sprintf("%s on line %d", p.Message, p.Found.Line)
}

func Parse(lexer *lex.Lexer) (ast.Node, error) {
	return ParseExpression(lexer)
}

func OperatorPrecedence(op string) uint8 {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "(", ")":
		return 3
	default:
		panic(fmt.Sprintf("Forgot OperatorPrecedence(%s) case", op))
	}
}

var EXPR_TOKEN_TYPES = []lex.TokenType{
	lex.TOKEN_TYPE_NUMBER, lex.TOKEN_TYPE_OP,
	lex.TOKEN_TYPE_RPAREN, lex.TOKEN_TYPE_LPAREN,
}

func nextToken(lexer *lex.Lexer, expected []lex.TokenType) (lex.Token, error) {
	token, err := lexer.Peek()
	if err != nil {
		return token, err
	}

	if !contains(token.TokenType, expected) {
		return token, Error{
			Found:    token,
			Expected: expected,
		}
	}

	return lexer.Next()
}

func ParseExpression(lexer *lex.Lexer) (ast.Node, error) {
	opStack, outQueue := []lex.Token{}, []lex.Token{}

	token, err := nextToken(lexer, EXPR_TOKEN_TYPES)
	if err != nil {
		return ast.Node{}, err
	}

	for {
		switch token.TokenType {
		case lex.TOKEN_TYPE_NUMBER:
			push(&outQueue, token)
		case lex.TOKEN_TYPE_OP:
			for len(opStack) > 0 && OperatorPrecedence(peekStack(opStack).Text) >= OperatorPrecedence(token.Text) && peekStack(opStack).Text != "(" {
				op := popStack(&opStack)
				push(&outQueue, op)
			}
			push(&opStack, token)
		case lex.TOKEN_TYPE_LPAREN:
			push(&opStack, token)
			fmt.Println("LENFTH OF OPSTACK", len(opStack))
		case lex.TOKEN_TYPE_RPAREN:
			for peekStack(opStack).Text != "(" {
				println(len(opStack), peekStack(opStack).Text)
				op := popStack(&opStack)
				push(&outQueue, op)
			}
			popStack(&opStack)
		default:
			panic("missing token type case expression")
		}

		token, err = lexer.Peek()
		if err != nil {
			return ast.Node{}, err
		}
		if !contains(token.TokenType, EXPR_TOKEN_TYPES) {
			break
		}
		token, _ = lexer.Next()
	}

	for len(opStack) > 0 {
		op := popStack(&opStack)
		push(&outQueue, op)
	}

	//for _, token := range outQueue {
	//	fmt.Printf("%s\n", token)
	//}

	//convert reverse polish notation into AST
	var nodeStack []ast.Node
	for _, token := range outQueue {
		switch token.TokenType {
		case lex.TOKEN_TYPE_NUMBER:
			nodeStack = append(nodeStack, ast.Node{
				NodeType: ast.NODE_TYPE_NUMBER,
				Children: nil,
				Token:    token,
			})
		case lex.TOKEN_TYPE_OP:
			//only binary operators right now
			if len(nodeStack) < 2 {
				switch token.Text {
				case "++", "--":
					child := nodeStack[len(nodeStack)-1]
					nodeStack = nodeStack[:len(nodeStack)-1]
					nodeStack = append(nodeStack, ast.Node{
						NodeType: ast.NODE_TYPE_UNOP,
						Children: []ast.Node{child},
						Token:    token,
					})
				}
			} else {
				//pop 2 off of end
				children := append([]ast.Node(nil), nodeStack[len(nodeStack)-2:]...)
				nodeStack = nodeStack[0 : len(nodeStack)-2]

				nodeStack = append(nodeStack, ast.Node{
					NodeType: ast.NODE_TYPE_BINOP,
					Children: children,
					Token:    token,
				})
			}
		default:
			panic("Unused case in convert opstack to node")
		}

	}

	if len(nodeStack) != 1 {
		return ast.Node{}, ExprError{
			Found:   outQueue[0],
			Message: "Parsed multiple expressions",
		}
	}

	return nodeStack[0], nil
}

func push[T lex.Token | ast.Node](collection *[]T, tok T) {
	*collection = append(*collection, tok)
}

func popStack[T lex.Token | ast.Node](st *[]T) T {
	if len(*st) == 0 {
		panic("pop on empty stack")
	}
	out := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return out
}

func peekStack[T lex.Token | ast.Node](st []T) T {
	return st[len(st)-1]
}

func popQueue(q *[]lex.Token) lex.Token {
	if len(*q) == 0 {
		panic("pop on empty stack")
	}
	out := (*q)[0]
	*q = (*q)[1:]
	return out
}

func peekQueue(q []lex.Token) lex.Token {
	return q[0]
}

func contains(ty lex.TokenType, slice []lex.TokenType) bool {
	for _, other := range slice {
		if ty == other {
			return true
		}
	}

	return false
}
