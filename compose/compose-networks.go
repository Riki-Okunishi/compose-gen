package compose

import (
	"fmt"
)

// networkEditor is
type networkEditor interface {
	String() string
}

// network describes one element of networks:
type network struct {
	name string
	version string
}

var _ networkEditor = &network{}

func newNetwork(n string, v string) networkEditor {
	return &network{
		name: n,
		version: v,
	}
}

func (n *network) String() string {
	// TODO: define String format
	str := fmt.Sprintf("%s%s:\n", indents[1], n.name)
	return str
}