package parse

import (
	"testing"
)

func TestSimpleExpression(t *testing.T) {
	//text := "1 * 2 + 3"
	//expected := []lex.Token{
	//	{TokenType: lex.TOKEN_TYPE_NUMBER, Text: "1"},
	//	{TokenType: lex.TOKEN_TYPE_NUMBER, Text: "2"},
	//	{TokenType: lex.TOKEN_TYPE_OP, Text: "*"},
	//	{TokenType: lex.TOKEN_TYPE_NUMBER, Text: "3"},
	//	{TokenType: lex.TOKEN_TYPE_OP, Text: "+"},
	//}
	//
	//buff := []lex.Token{}
	//lexer := lex.NewLexer(lex.NewSource(text))
	//
	//for tmp, _ := lexer.Next(); tmp.TokenType != lex.TOKEN_TYPE_EOF; tmp, _ = lexer.Next() {
	//	buff = append(buff, tmp)
	//}
	//
	//result := ParseExpression(buff)
	//for i := range result {
	//	if result[i].TokenType != expected[i].TokenType || result[i].Text != expected[i].Text {
	//		t.Fatal("unexpected token", result[i])
	//	}
	//}
	panic("TODO")
}
