package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type YieldFrom struct {
	position *node.Position
	Expr     node.Node
}

func NewYieldFrom(Expression node.Node) node.Node {
	return &YieldFrom{
		nil,
		Expression,
	}
}

func (n YieldFrom) Attributes() map[string]interface{} {
	return nil
}

func (n YieldFrom) Position() *node.Position {
	return n.position
}

func (n YieldFrom) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n YieldFrom) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
