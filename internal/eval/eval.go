package eval

import (
	"awesomeProject/internal/ast"
	"fmt"
	"strconv"
)

type eval struct {
	result int
}

func (e *eval) VisitNumber(parent *ast.Node, number *ast.Node) {
	result, err := strconv.ParseInt(number.Token.Text, 10, 32)
	if err != nil {
		panic(err)
	}

	e.result = int(result)
}

func (e *eval) VisitBinop(parent *ast.Node, op *ast.Node) {
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

func (e *eval) VisitUnop(parent *ast.Node, op *ast.Node) {
	child := EvalExpression(&op.Children[0])
	fmt.Println("HERE")
	switch op.Token.Text {
	case "++":
		e.result = child + 1
	case "--":
		e.result = child - 1
	}
}

func EvalExpression(expr *ast.Node) int {
	var eval eval
	expr.Visit(nil, &eval)

	return eval.result
}
