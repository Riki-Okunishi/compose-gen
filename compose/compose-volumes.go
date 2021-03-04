package compose

import (
	"fmt"
)

// volumes is map of volume
type volumeEditor interface {
}

// volume describes one element of volumes:
type volume struct {
	name string
	version string
}

func newVolume(n string, v string) volumeEditor {
	return &volume{
		name: n,
		version: v,
	}
}

func (v *volume) String() string {
	// TODO: define Stringer format
	return fmt.Sprintf("")
}