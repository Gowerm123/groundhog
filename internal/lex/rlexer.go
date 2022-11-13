package lex

import (
	"reflect"
	"strconv"
)

var breakingSyms []rune = []rune{' ', '"', '\'', '(', ')'}

func (lexer *Lexer) RNext() (Token, error) {
	if lexer.source.Peek() == ' ' {
		for lexer.source.Peek() == ' ' {
			lexer.source.Next()
		}
	}
	return lexer.nextHelper(nil)
}

func (lexer *Lexer) nextHelper(buff []rune) (Token, error) {
	if contains(breakingSyms, lexer.source.Peek()) {
		char := lexer.source.Next()
		return lexer.toToken(append(buff, char))
	}
	buff = append(buff, lexer.source.Next())
	return lexer.nextHelper(buff)
}

func (lexer *Lexer) toToken(buff []rune) (Token, error) {
	switch string(buff) {
	case "+", "-", "*", "/", "++", "+=", "-=", "--", "**":
		return Token{
			TokenType: TOKEN_TYPE_OP,
			Col:       lexer.source.col,
			Line:      lexer.source.line,
			Text:      string(lexer.source.Next()),
		}, nil
	case "(":
		lexer.source.Next()
		return Token{
			TokenType: TOKEN_TYPE_LPAREN,
			Text:      "(",
			Col:       lexer.source.col,
			Line:      lexer.source.line,
		}, nil
	case ")":
		lexer.source.Next()
		return Token{
			TokenType: TOKEN_TYPE_RPAREN,
			Text:      ")",
			Col:       lexer.source.col,
			Line:      lexer.source.line,
		}, nil
	default:
		//is a number, refactor this later
		if _, err := strconv.Atoi(string(buff)); err == nil {
			return Token{
				TokenType: TOKEN_TYPE_NUMBER,
				Text:      string(buff),
				Col:       lexer.source.col,
				Line:      lexer.source.line,
			}, nil
		}
		return Token{}, Error{
			Col:      lexer.source.col,
			Line:     lexer.source.line,
			Found:    string(lexer.source.Peek()),
			Expected: "expression",
		}
	}
}

func contains[T any](collection []T, target T) bool {
	for _, t := range collection {
		if reflect.DeepEqual(t, target) {
			return true
		}
	}

	return false
}
