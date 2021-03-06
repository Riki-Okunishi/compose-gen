package compose

import (
	"fmt"
)

// volumes is map of volume
type volumeEditor interface {
	String() string
}

// volume describes one element of volumes:
type volume struct {
	name string
	version string
}

var _ volumeEditor = &volume{}

func newVolume(n string, v string) volumeEditor {
	return &volume{
		name: n,
		version: v,
	}
}

func (v *volume) String() string {
	// TODO: define String format
	str := fmt.Sprintf("%s%s:\n", indents[1], v.name)
	return str
}