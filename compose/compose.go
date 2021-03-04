package compose

import (
	"fmt"
)

const defaultVersion string = "3.5"

// ComposeFileEditor is model of docker-compose.yml in specific version
type ComposeFileEditor interface {
	GetVersion() string
	Services(srv string) serviceEditor
	Networks(nw string) networkEditor
	Volumes(vol string) volumeEditor
}

type composeFile struct {
	version  string
	services map[string]serviceEditor
	networks map[string]networkEditor
	volumes  map[string]volumeEditor
}

// NewDefaultComposeFile returns composeFile object in default version
func NewDefaultComposeFile() ComposeFileEditor {
	return NewComposeFile(defaultVersion)
}

// NewComposeFile returns composeFile object in given version
func NewComposeFile(v string) ComposeFileEditor {
	// TODO: validate whether version is supported version or not
	return &composeFile{
		version: v,
		services: map[string]serviceEditor{},
		networks: map[string]networkEditor{},
		volumes: map[string]volumeEditor{},
	}
}

func (c *composeFile) String() string {
	// TODO: define Stringer format
	return fmt.Sprintf("version: %s\n", c.version)
}

// GetVersion returns this compose file version
func (c *composeFile) GetVersion() string {
	return c.version
}

// Services returns Service object corresponding to service. If not exist, add it into services.
func (c *composeFile) Services(srv string) serviceEditor {
	if s, ok := c.services[srv]; ok {
		return s
	}
	s := newService(srv, c.version)
	c.services[srv] = s
	return s
}

// Networks returns Network object corresponding to network. If not exist, add it into networks.
func (c *composeFile) Networks(nw string) networkEditor {
	if n, ok := c.networks[nw]; ok {
		return n
	}
	n := newNetwork(nw, c.version)
	c.networks[nw] = n
	return n
}

// Volumes returns Volume object corresponding to volume. If not exist, add it into volumes.
func (c *composeFile) Volumes(vol string) volumeEditor {
	if v, ok := c.volumes[vol]; ok {
		return v
	}
	v := newVolume(vol, c.version)
	c.volumes[vol] = v
	return v
}
