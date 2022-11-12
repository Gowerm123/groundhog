package ast

import "awesomeProject/internal/tokens"

type NodeType uint8

const (
	NODE_TYPE_NUMBER = iota
	NODE_TYPE_BINOP
)

type Node struct {
	nodeType NodeType
	Children []Node
	Left     *Node
	Right    *Node
	Parent   *Node
	toks     []tokens.Token
}

func NewAstNode(nodeType NodeType, toks []tokens.Token) Node {
	return Node{
		nodeType: nodeType,
		toks:     toks,
	}
}
