package eval

import (
	"awesomeProject/internal/parse"
	"fmt"
	"strconv"
)

type eval struct {
	result int
}

func (e *eval) VisitNumber(parent *parse.Node, number *parse.Node) {
	result, err := strconv.ParseInt(number.Token.Text, 10, 32)
	if err != nil {
		panic(err)
	}

	e.result = int(result)
}

func (e *eval) VisitBinop(parent *parse.Node, op *parse.Node) {
	left := EvalExpression(&op.Children[0])
	right := EvalExpression(&op.Children[1])

	switch op.Token.Text {
	case "+":
		e.result = left + right
	case "-":
		e.result = left - right
	case "*":
		e.result = left * right
	case "/":
		e.result = left / right
	default:
		panic(fmt.Sprintf("Don't know how to eval op (%s)", op.Token.Text))
	}
}

func EvalExpression(expr *parse.Node) int {
	var eval eval
	expr.Visit(nil, &eval)

	return eval.result
}