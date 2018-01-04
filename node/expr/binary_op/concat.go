package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Concat struct {
	BinaryOp
}

func NewConcat(Variable node.Node, Expression node.Node) node.Node {
	return &Concat{
		BinaryOp{
			nil,
			Variable,
			Expression,
		},
	}
}

func (n Concat) Attributes() map[string]interface{} {
	return nil
}

func (n Concat) Position() *node.Position {
	return n.position
}

func (n Concat) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Concat) Walk(v node.Visitor) {
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
