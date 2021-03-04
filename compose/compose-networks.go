package compose

import (
	"fmt"
)

// networkEditor is
type networkEditor interface {

}

// network describes one element of networks:
type network struct {
	name string
	version string
}

func newNetwork(n string, v string) networkEditor {
	return &network{
		name: n,
		version: v,
	}
}

func (n *network) String() string {
	// TODO: define Stringer format
	return fmt.Sprintf("")
}