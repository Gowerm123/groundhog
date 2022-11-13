package printer

import (
	"awesomeProject/internal/ast"
	"awesomeProject/internal/eval"
	"awesomeProject/internal/lex"
	"awesomeProject/internal/parse"
	"testing"
)

func doParse(text string, t *testing.T) ast.Node {
	lexer := lex.NewLexer(lex.NewSource(text))
	node, err := parse.Parse(&lexer)
	if err != nil {
		t.Fatal(err)
	}

	return node
}

func doTest(text string, t *testing.T) {
	// test by cycling printed output into compiler again
	// assert same text and same result
	expectedNode := doParse(text, t)
	expectedText := String(&expectedNode)
	expectedResult := eval.EvalExpression(&expectedNode)

	foundNode := doParse(expectedText, t)
	foundText := String(&foundNode)
	foundResult := eval.EvalExpression(&foundNode)

	if expectedText != foundText || expectedResult != foundResult {
		t.Fatalf(
			"Printed expression did produce identical compiler result!\nOrig: %s\nPrinted: %s=%d\nCycled: %s=%d",
			text,
			expectedText, expectedResult,
			foundText, foundResult,
		)
	}
}

func TestExpr1(t *testing.T) {
	doTest("3 + 2 * 1", t)

}

func TestExpr2(t *testing.T) {
	doTest("1 * 2 + 3", t)
}

func TestExpr3(t *testing.T) {
	doTest("1 + 21 *3/4 - (10+ 15)", t)
}
