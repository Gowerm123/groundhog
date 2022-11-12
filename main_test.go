package main

import (
	"fmt"
	"testing"
)

func TestLexExpr1(t *testing.T) {
	lexer := NewLexer(NewSource("1 + 2 *3/4 - (10+ 15)"))
	expected := []Token{
		{TokenType: TOKEN_TYPE_NUMBER, Text: "1"},
		{TokenType: TOKEN_TYPE_OP, Text: "+"},
		{TokenType: TOKEN_TYPE_NUMBER, Text: "2"},
		{TokenType: TOKEN_TYPE_OP, Text: "*"},
		{TokenType: TOKEN_TYPE_NUMBER, Text: "3"},
		{TokenType: TOKEN_TYPE_OP, Text: "/"},
		{TokenType: TOKEN_TYPE_NUMBER, Text: "4"},
		{TokenType: TOKEN_TYPE_OP, Text: "-"},
		{TokenType: TOKEN_TYPE_START_EXPR},
		{TokenType: TOKEN_TYPE_NUMBER, Text: "10"},
		{TokenType: TOKEN_TYPE_OP, Text: "+"},
		{TokenType: TOKEN_TYPE_NUMBER, Text: "15"},
		{TokenType: TOKEN_TYPE_END_EXPR},
		{TokenType: TOKEN_TYPE_EOF},
	}

	for _, check := range expected {
		token, err := lexer.Next()
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		fmt.Printf("%s\n", token)
		if token.TokenType != check.TokenType || token.Text != check.Text {
			t.Fatalf("%s != %s or '%s' != '%s'", token.TokenType, check.TokenType, token.Text, check.Text)
		}
	}
}

func TestReadChars(t *testing.T) {
	source := NewSource("hello")

	if source.Peek() != 'h' {
		t.FailNow()
	}
	if source.Peek() != 'h' {
		t.FailNow()
	}
	if source.Next() != 'h' {
		t.FailNow()
	}
	if source.Next() != 'e' {
		t.FailNow()
	}
	if source.Next() != 'l' {
		t.FailNow()
	}
	if source.Peek() != 'l' {
		t.FailNow()
	}
	if source.Peek() != 'l' {
		t.FailNow()
	}
	if source.Next() != 'l' {
		t.FailNow()
	}
	if source.Next() != 'o' {
		t.FailNow()
	}
	if source.Next() != -1 {
		t.FailNow()
	}
	if source.Next() != -1 {
		t.FailNow()
	}
	if source.Peek() != -1 {
		t.FailNow()
	}
	if source.Peek() != -1 {
		t.FailNow()
	}
	if source.Next() != -1 {
		t.FailNow()
	}
}
