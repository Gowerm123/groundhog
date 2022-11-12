package main

import (
	"awesomeProject/internal/lexer"
	"awesomeProject/internal/parser"
	"awesomeProject/internal/tokens"
	"fmt"
)

func main() {
	text := "1 * 2 + 3"
	lexer := lexer.NewLexer(lexer.NewSource(text))
	exprBuff := []tokens.Token{}
	for token, err := lexer.Next(); token.TokenType != tokens.TOKEN_TYPE_EOF && err == nil; token, err = lexer.Next() {
		exprBuff = append(exprBuff, token)
	}

	parsed := parser.ParseExpression(exprBuff)

	for _, k := range parsed {
		fmt.Println(k)
	}
}
