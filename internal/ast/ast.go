package ast

import (
	"awesomeProject/internal/lex"
	"fmt"
)

type NodeType uint8

const (
	NODE_TYPE_NUMBER = iota
	NODE_TYPE_BINOP
	NODE_TYPE_UNOP
)

type Node struct {
	NodeType NodeType
	Children []Node
	Token    lex.Token
}

type NodeVisitor interface {
	VisitNumber(parent *Node, number *Node)
	VisitBinop(parent *Node, op *Node)
	VisitUnop(parent *Node, op *Node)
}

func (n *Node) Visit(parent *Node, visitor NodeVisitor) {
	switch n.NodeType {
	case NODE_TYPE_NUMBER:
		visitor.VisitNumber(parent, n)
	case NODE_TYPE_BINOP:
		visitor.VisitBinop(parent, n)
	case NODE_TYPE_UNOP:
		visitor.VisitUnop(parent, n)
	default:
		panic(fmt.Sprintf("Node.Visit(...) => Forgot node type case %d", n.NodeType))
	}
}
