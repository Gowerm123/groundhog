package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Source struct {
	col      uint
	line     uint
	nextChar rune
	index    int
	chars    []rune
}

func NewSource(source string) Source {
	return Source{
		col:      1,
		line:     1,
		nextChar: 0,
		index:    0,
		chars:    []rune(source),
	}
}

func (source *Source) Next() rune {
	if source.nextChar != 0 {
		ch := source.nextChar
		source.nextChar = 0
		return ch
	}

	if source.index >= len(source.chars) {
		return -1
	}

	ch := source.chars[source.index]
	source.index += 1
	if ch == '\n' {
		source.col = 1
		source.line += 1
	} else {
		source.col += 1
	}

	return ch
}

func (source *Source) Peek() rune {
	if source.nextChar == 0 {
		source.nextChar = source.Next()
	}

	return source.nextChar
}

func (source *Source) SkipWhitespace() {
	for unicode.IsSpace(source.Peek()) {
		source.Next()
	}
}

type TokenType uint8

//1 + 2 * 3 / 44 - (10 + 15)
const (
	TOKEN_TYPE_EOF = iota
	TOKEN_TYPE_NUMBER
	TOKEN_TYPE_OP
	TOKEN_TYPE_START_EXPR
	TOKEN_TYPE_END_EXPR
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

type LexError struct {
	Found    string
	Expected string
	Col      uint
	Line     uint
}

func (lexError LexError) Error() string {
	return fmt.Sprintf("Found '%s' at line %d, column %d, but expected one of [%s]", lexError.Found, lexError.Line, lexError.Col, lexError.Expected)
}

type Lexer struct {
	source      Source
	peekedToken Token
}

func NewLexer(source Source) Lexer {
	return Lexer{
		source: source,
		peekedToken: Token{
			TokenType: TOKEN_TYPE_EOF,
		},
	}
}

func (lexer *Lexer) Peek() (Token, error) {
	if lexer.peekedToken.TokenType == TOKEN_TYPE_EOF {
		token, err := lexer.Next()
		if err != nil {
			return Token{}, err
		}

		lexer.peekedToken = token
	}

	return lexer.peekedToken, nil
}

func (lexer *Lexer) nextNumber() (Token, error) {
	var str strings.Builder
	token := Token{
		TokenType: TOKEN_TYPE_NUMBER,
		Col:       lexer.source.col,
		Line:      lexer.source.line,
	}

	digit := lexer.source.Peek()
	for digit >= '0' && digit <= '9' {
		str.WriteRune(lexer.source.Next())
		digit = lexer.source.Peek()
	}

	token.Text = str.String()
	return token, nil
}

func (lexer *Lexer) Next() (Token, error) {
	if lexer.peekedToken.TokenType != TOKEN_TYPE_EOF {
		token := lexer.peekedToken
		lexer.peekedToken.TokenType = TOKEN_TYPE_EOF

		return token, nil
	}

	lexer.source.SkipWhitespace()
	switch lexer.source.Peek() {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return lexer.nextNumber()
	case '+', '-', '*', '/':
		return Token{
			TokenType: TOKEN_TYPE_OP,
			Col:       lexer.source.col,
			Line:      lexer.source.line,
			Text:      string(lexer.source.Next()),
		}, nil
	case -1:
		return Token{
			TokenType: TOKEN_TYPE_EOF,
			Col:       lexer.source.col,
			Line:      lexer.source.line,
		}, nil
	case '(':
		lexer.source.Next()
		return Token{
			TokenType: TOKEN_TYPE_START_EXPR,
			Col:       lexer.source.col,
			Line:      lexer.source.line,
		}, nil
	case ')':
		lexer.source.Next()
		return Token{
			TokenType: TOKEN_TYPE_END_EXPR,
			Col:       lexer.source.col,
			Line:      lexer.source.line,
		}, nil
	default:
		return Token{}, LexError{
			Col:      lexer.source.col,
			Line:     lexer.source.line,
			Found:    string(lexer.source.Peek()),
			Expected: "expression",
		}
	}
}

func main() {

}
