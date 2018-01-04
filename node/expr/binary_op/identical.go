package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Identical struct {
	BinaryOp
}

func NewIdentical(Variable node.Node, Expression node.Node) node.Node {
	return &Identical{
		BinaryOp{
			nil,
			Variable,
			Expression,
		},
	}
}

func (n Identical) Attributes() map[string]interface{} {
	return nil
}

func (n Identical) Position() *node.Position {
	return n.position
}

func (n Identical) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Identical) Walk(v node.Visitor) {
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
