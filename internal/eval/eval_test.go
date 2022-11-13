package eval

import (
	"awesomeProject/internal/lex"
	"awesomeProject/internal/parse"
	"fmt"
	"testing"
)

func doEval(text string, t *testing.T) float64 {
	lexer := lex.NewLexer(lex.NewSource(text))
	node, err := parse.Parse(&lexer)
	if err != nil {
		t.Fatal(err)
	}

	return EvalExpression(&node)
}

func TestExpr1(t *testing.T) {
	text := "4 + 3 * 2 / 2"
	expected := 4 + 3*2/2
	found := doEval(text, t)
	if expected != int(found) {
		t.Fatalf("\n%s == %d\n%s != %f", text, expected, text, found)
	}
}

func TestExpr2(t *testing.T) {
	text := "2 / 2 * 3 + 4"
	expected := 2/2*3 + 4
	found := doEval(text, t)
	if expected != int(found) {
		t.Fatalf("\n%s == %d\n%s != %f", text, expected, text, found)
	}
}

func TestExpr3(t *testing.T) {
	text := "1 + 21 *3/4 - (10+ 15)"
	expected := 1 + 21*3/4 - (10 + 15)
	found := doEval(text, t)
	if expected != int(found) {
		t.Fatalf("\n%s == %d\n%s != %f", text, expected, text, found)
	}
}

func TestExpr4(t *testing.T) {
	text := "1.0 + 3.5"
	expected := 4.5
	found := doEval(text, t)
	if expected != found {
		t.Fatalf("\n%s == %f\n%s != %f", text, expected, text, found)
	}
}

func TestExpr5(t *testing.T) {
	text := "1 + 2.5*(3 - 9.2)"
	expected := 1 + 2.5*(3-9.2)
	found := doEval(text, t)
	if expected != found {
		fmt.Println(expected, found)
		t.Fatalf("\n%s == %f\n%s != %f", text, expected, text, found)
	}
}
