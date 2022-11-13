package printer

import (
	"awesomeProject/internal/ast"
	"strings"
)

type stringer struct {
	buf strings.Builder
}

func (s *stringer) VisitNumber(parent *ast.Node, number *ast.Node) {
	s.buf.WriteString(number.Token.Text)
}

func (s *stringer) VisitBinop(parent *ast.Node, op *ast.Node) {
	s.buf.WriteRune('(')
	op.Children[0].Visit(op, s)
	s.buf.WriteRune(' ')
	s.buf.WriteString(op.Token.Text)
	s.buf.WriteRune(' ')
	op.Children[1].Visit(op, s)
	s.buf.WriteRune(')')
}

func (s *stringer) VisitUnop(parent *ast.Node, op *ast.Node) {

}

func String(node *ast.Node) string {
	var toString stringer
	node.Visit(nil, &toString)

	return toString.buf.String()
}
