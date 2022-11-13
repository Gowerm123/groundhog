package parse

import (
	"awesomeProject/internal/ast"
	"awesomeProject/internal/lex"
	"awesomeProject/internal/printer"
	"testing"
)

func number(text string) ast.Node {
	return ast.Node{
		NodeType: ast.NODE_TYPE_NUMBER,
		Token: lex.Token{
			TokenType: lex.TOKEN_TYPE_NUMBER,
			Text:      text,
		},
	}
}

func binop(op string, left ast.Node, right ast.Node) ast.Node {
	return ast.Node{
		NodeType: ast.NODE_TYPE_BINOP,
		Token: lex.Token{
			TokenType: lex.TOKEN_TYPE_OP,
			Text:      op,
		},
		Children: []ast.Node{left, right},
	}
}

func parse(text string, t *testing.T) ast.Node {
	source := lex.NewSource(text)
	lexer := lex.NewLexer(source)
	node, err := Parse(&lexer)
	if err != nil {
		t.Fatal(err)
	}

	return node
}

func assertSame(left ast.Node, right ast.Node, t *testing.T) {
	//todo: maybe put some info here to let you know "what" isn't the same?
	//      would have to modify the helper functions to keep track of col/line

	if left.NodeType != right.NodeType || left.Token.Text != right.Token.Text || left.Token.TokenType != right.Token.TokenType || len(left.Children) != len(right.Children) {
		t.Fatalf("Nodes are not equal:\n%s\n%s\n", printer.String(&left), printer.String(&right))
	}

	for i := range left.Children {
		assertSame(left.Children[i], right.Children[i], t)
	}
}

func TestExpr1(t *testing.T) {
	text := "3 + 2 * 1"
	expected := binop(
		"+",
		number("3"),
		binop("*", number("2"), number("1")),
	)
	found := parse(text, t)

	assertSame(found, expected, t)
}

func TestExpr2(t *testing.T) {
	text := "1 * 2 + 3"
	expected := binop(
		"+",
		binop("*", number("1"), number("2")),
		number("3"),
	)
	found := parse(text, t)

	assertSame(found, expected, t)
}

func TestExpr3(t *testing.T) {
	text := "1 + 21 *3/4 - (10+ 15)"
	expected := binop(
		"-",
		binop("+", number("1"), binop("/", binop("*", number("21"), number("3")), number("4"))),
		binop("+", number("10"), number("15")),
	)
	found := parse(text, t)

	assertSame(found, expected, t)
}
