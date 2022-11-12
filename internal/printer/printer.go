package printer

import (
	"awesomeProject/internal/parse"
	"strings"
)

type stringer struct {
	buf strings.Builder
}

func (s *stringer) VisitNumber(parent *parse.Node, number *parse.Node) {
	s.buf.WriteString(number.Token.Text)
}

func (s *stringer) VisitBinop(parent *parse.Node, op *parse.Node) {
	s.buf.WriteRune('(')
	op.Children[0].Visit(op, s)
	s.buf.WriteRune(' ')
	s.buf.WriteString(op.Token.Text)
	s.buf.WriteRune(' ')
	op.Children[1].Visit(op, s)
	s.buf.WriteRune(')')
}

func String(node *parse.Node) string {
	var toString stringer
	node.Visit(nil, &toString)

	return toString.buf.String()
}
