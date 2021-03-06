package compose

import (
	"fmt"
	"os"
)

const (
	defaultVersion string = "3.5"
)

var (
	indents = []string{
		"",
		"  ",
		"    ",
		"      ",
		"        ",
	}
)

// ComposeFileEditor is model of docker-compose.yml in specific version
type ComposeFileEditor interface {
	String() string
	GenerateFile(path string) error
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

type valueSyntax interface {
	GetSyntaxType() string
	String() string
}

var _ ComposeFileEditor = &composeFile{}

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

func (c *composeFile) String() string {
	// TODO: define String format
	str := fmt.Sprintf("version: \"%s\"\n", c.version)
	if len(c.services) != 0 {
		str += "services:\n"
		for _, s := range c.services {
			str += s.String()
		}
	}
	if len(c.networks) != 0 {
		str += "networks:\n"
		for _, n := range c.networks {
			str += n.String()
		}
	}
	if len(c.volumes) != 0 {
		str += "volumes:\n"
		for _, v := range c.volumes {
			str += v.String()
		}
	}
	return str
}

func (c *composeFile) GenerateFile(path string) error {
	fp, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fp.Close()

	fp.Write([]byte(c.String()))
	return nil
}