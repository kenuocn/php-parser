package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Label struct {
	position  *node.Position
	LabelName node.Node
}

func NewLabel(LabelName node.Node) node.Node {
	return &Label{
		nil,
		LabelName,
	}
}

func (n Label) Attributes() map[string]interface{} {
	return nil
}

func (n Label) Position() *node.Position {
	return n.position
}

func (n Label) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Label) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.LabelName != nil {
		vv := v.GetChildrenVisitor("LabelName")
		n.LabelName.Walk(vv)
	}

	v.LeaveNode(n)
}
