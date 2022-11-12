package parser

import (
	"awesomeProject/internal/lexer"
	"awesomeProject/internal/tokens"
	"testing"
)

func TestSimpleExpression(t *testing.T) {
	text := "1 * 2 + 3"
	expected := []tokens.Token{
		{TokenType: tokens.TOKEN_TYPE_NUMBER, Text: "1"},
		{TokenType: tokens.TOKEN_TYPE_NUMBER, Text: "2"},
		{TokenType: tokens.TOKEN_TYPE_OP, Text: "*"},
		{TokenType: tokens.TOKEN_TYPE_NUMBER, Text: "3"},
		{TokenType: tokens.TOKEN_TYPE_OP, Text: "+"},
	}

	buff := []tokens.Token{}
	lexer := lexer.NewLexer(lexer.NewSource(text))

	for tmp, _ := lexer.Next(); tmp.TokenType != tokens.TOKEN_TYPE_EOF; tmp, _ = lexer.Next() {
		buff = append(buff, tmp)
	}

	result := ParseExpression(buff)
	for i := range result {
		if result[i].TokenType != expected[i].TokenType || result[i].Text != expected[i].Text {
			t.Fatal("unexpected token", result[i])
		}
	}
}
