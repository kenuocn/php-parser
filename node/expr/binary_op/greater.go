package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Greater struct {
	BinaryOp
}

func NewGreater(Variable node.Node, Expression node.Node) node.Node {
	return &Greater{
		BinaryOp{
			nil,
			Variable,
			Expression,
		},
	}
}

func (n Greater) Attributes() map[string]interface{} {
	return nil
}

func (n Greater) Position() *node.Position {
	return n.position
}

func (n Greater) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Greater) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		vv := v.GetChildrenVisitor("Left")
		n.Left.Walk(vv)
	}

	if n.Right != nil {
		vv := v.GetChildrenVisitor("Right")
		n.Right.Walk(vv)
	}

	v.LeaveNode(n)
}
