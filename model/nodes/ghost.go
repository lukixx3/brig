package nodes

import (
	"fmt"

	capnp_model "github.com/disorganizer/brig/model/nodes/capnp"
	capnp "zombiezen.com/go/capnproto2"
)

type Ghost struct {
	Node

	oldType NodeType
}

// MakeGhost takes an existing node and converts it to a ghost.
// In the ghost form no metadata is lost, but the node should
// not show up.
func MakeGhost(nd Node) (*Ghost, error) {
	return &Ghost{
		Node:    nd,
		oldType: nd.Type(),
	}, nil
}

func (g *Ghost) Type() NodeType {
	return NodeTypeGhost
}

func (g *Ghost) ToCapnp() (*capnp.Message, error) {
	oldMsg, err := g.Node.ToCapnp()
	if err != nil {
		return nil, err
	}

	oldData, err := oldMsg.Marshal()
	if err != nil {
		return nil, err
	}

	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	capghost, err := capnp_model.NewRootGhost(seg)
	if err != nil {
		return nil, err
	}

	if err := capghost.SetOldNode(oldData); err != nil {
		return nil, err
	}

	capghost.SetOldType(uint8(g.oldType))
	return msg, nil
}

func (g *Ghost) FromCapnp(msg *capnp.Message) error {
	capghost, err := capnp_model.ReadRootGhost(msg)
	if err != nil {
		return err
	}

	// Make sure g.Node is initialized with a correct struct.
	switch typ := capghost.OldType(); typ {
	case NodeTypeCommit:
		g.Node = &Commit{}
	case NodeTypeDirectory:
		g.Node = &Directory{}
	case NodeTypeFile:
		g.Node = &File{}
	default:
		panic(fmt.Sprintf("Unsupported node type: %v", typ))
	}

	oldData, err := capghost.OldNode()
	if err != nil {
		return err
	}

	oldMsg, err := capnp.Unmarshal(oldData)
	if err != nil {
		return err
	}

	return g.Node.FromCapnp(oldMsg)
}
