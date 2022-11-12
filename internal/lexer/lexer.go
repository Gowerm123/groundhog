package lexer

import (
	"awesomeProject/internal/tokens"
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
	peekedToken tokens.Token
}

func NewLexer(source Source) Lexer {
	return Lexer{
		source: source,
		peekedToken: tokens.Token{
			TokenType: tokens.TOKEN_TYPE_EOF,
		},
	}
}

func (lexer *Lexer) Peek() (tokens.Token, error) {
	if lexer.peekedToken.TokenType == tokens.TOKEN_TYPE_EOF {
		token, err := lexer.Next()
		if err != nil {
			return tokens.Token{}, err
		}

		lexer.peekedToken = token
	}

	return lexer.peekedToken, nil
}

func (lexer *Lexer) nextNumber() (tokens.Token, error) {
	var str strings.Builder
	token := tokens.Token{
		TokenType: tokens.TOKEN_TYPE_NUMBER,
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

func (lexer *Lexer) Next() (tokens.Token, error) {
	if lexer.peekedToken.TokenType != tokens.TOKEN_TYPE_EOF {
		token := lexer.peekedToken
		lexer.peekedToken.TokenType = tokens.TOKEN_TYPE_EOF

		return token, nil
	}

	lexer.source.SkipWhitespace()
	switch lexer.source.Peek() {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return lexer.nextNumber()
	case '+', '-', '*', '/':
		return tokens.Token{
			TokenType: tokens.TOKEN_TYPE_OP,
			Col:       lexer.source.col,
			Line:      lexer.source.line,
			Text:      string(lexer.source.Next()),
		}, nil
	case -1:
		return tokens.Token{
			TokenType: tokens.TOKEN_TYPE_EOF,
			Col:       lexer.source.col,
			Line:      lexer.source.line,
		}, nil
	case '(':
		lexer.source.Next()
		return tokens.Token{
			TokenType: tokens.TOKEN_TYPE_START_EXPR,
			Col:       lexer.source.col,
			Line:      lexer.source.line,
		}, nil
	case ')':
		lexer.source.Next()
		return tokens.Token{
			TokenType: tokens.TOKEN_TYPE_END_EXPR,
			Col:       lexer.source.col,
			Line:      lexer.source.line,
		}, nil
	default:
		return tokens.Token{}, LexError{
			Col:      lexer.source.col,
			Line:     lexer.source.line,
			Found:    string(lexer.source.Peek()),
			Expected: "expression",
		}
	}
}
