package eval

import (
	"awesomeProject/internal/ast"
	"fmt"
	"math"
	"strconv"
)

type eval struct {
	result float64
}

func (e *eval) VisitNumber(parent *ast.Node, number *ast.Node) {
	result, err := strconv.ParseFloat(number.Token.Text, 64)
	if err != nil {
		panic(err)
	}

	e.result = result
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
	case "**":
		e.result = pow(left, right)
	default:
		panic(fmt.Sprintf("Don't know how to eval op (%s)", op.Token.Text))
	}
}

func pow[T int | float64](base T, pow T) T {
	val := math.Pow(float64(base), float64(pow))

	return T(val)
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

func EvalExpression(expr *ast.Node) float64 {
	var eval eval
	expr.Visit(nil, &eval)

	return eval.result
}
