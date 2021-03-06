package compose

import (
	"fmt"
)

// services is map of service
type serviceEditor interface {
	String() string
	Build(b string) error
	Image(i string) error
	Ports(arg interface{}, args ...interface{}) error
	Volumes(vs ...string) error
	DependsOn(ds ...string) error
}

// service describes one element of services:
type service struct {
	name      string
	version   string
	build     string
	image     string
	ports     []valueSyntax
	volumes   []string
	dependsOn []string
}

type portsLongSyntax struct {
	syntaxType string
	target     uint
	published  uint
	protocol   string
	mode       string
}

type portsShortSyntax struct {
	syntaxType string
	ports      string
}

var (
	_ serviceEditor = &service{}
	_ valueSyntax   = &portsLongSyntax{}
	_ valueSyntax   = &portsShortSyntax{}
)

func newService(n string, v string) serviceEditor {
	return &service{
		name:      n,
		version:   v,
		build:     "",
		image:     "",
		ports:     []valueSyntax{},
		volumes:   []string{},
		dependsOn: []string{},
	}
}

/*func (s *service) GetName() string {
	return s.name
}*/

func (s *service) Build(b string) error {
	// TODO: Validate whether b is dir path
	if s.image != "" {
		return fmt.Errorf("Error: service %s is already allocated image: %s", s.name, s.image)
	}
	s.build = b
	return nil
}

func (s *service) Image(i string) error {
	// TODO: Validate whether image exist here or Docker Hub
	if s.build != "" {
		return fmt.Errorf("Error: service %s is already allocated build: %s", s.name, s.build)
	}
	s.image = i
	return nil
}

func (s *service) Ports(arg interface{}, args ...interface{}) error {
	switch v := arg.(type) {
	case string:
		pss, err := newPortsShortSyntax(v)
		if err != nil {
			return err
		}
		s.ports = append(s.ports, pss)
	case []string:
		if len(args) > 0 {
			return fmt.Errorf("Error: invalid argument args %v", args)
		}
		for _, p := range v {
			pss, err := newPortsShortSyntax(p)
			if err != nil {
				return err
			}
			s.ports = append(s.ports, pss)
		}
	case map[string]interface{}:
		pls, err := newPortsLongSyntax(v)
		if err != nil {
			return err
		}
		s.ports = append(s.ports, pls)

	case []map[string]interface{}:
		if len(args) > 0 {
			return fmt.Errorf("Error: invalid argument args %v", args)
		}
		for _, m := range v {
			pls, err := newPortsLongSyntax(m)
			if err != nil {
				return err
			}
			s.ports = append(s.ports, pls)
		}

	default:
		return fmt.Errorf("Error: invalid argument arg %v", arg)
	}
	return nil
}

func (s *service) Volumes(vs ...string) error {
	for _, v := range vs {
		// TODO: Validate whether all of vs is dir path
		s.volumes = append(s.volumes, v)
	}
	return nil
}

func (s *service) DependsOn(ds ...string) error {
	for _, d := range ds {
		// Note: services included in dependsOn need to be allocated in ComposeFile.services as theirs same name
		s.dependsOn = append(s.dependsOn, d)
	}
	return nil
}

func (s *service) String() string {
	// TODO: define Stringer format
	str := indents[1] + s.name + ":\n"
	if s.build != "" && s.image == "" {
		str += fmt.Sprintf("%sbuild: %s\n", indents[2], s.build)
	}
	if s.image != "" && s.build == "" {
		str += fmt.Sprintf("%simage: %s\n", indents[2], s.image)
	}
	if len(s.ports) != 0 {
		str += fmt.Sprintf("%sports:\n", indents[2])
		for _, v := range s.ports {
			str += fmt.Sprintf("%s- %v", indents[3], v)
		}
	}
	if len(s.volumes) != 0 {
		str += fmt.Sprintf("%svolumes:\n", indents[2])
		for _, v := range s.volumes {
			str += fmt.Sprintf("%s- %v\n", indents[3], v)
		}
	}
	if len(s.dependsOn) != 0 {
		str += fmt.Sprintf("%sdepends_on:\n", indents[2])
		for _, d := range s.dependsOn {
			str += fmt.Sprintf("%s- %v\n", indents[3], d)
		}
	}
	return str
}

func newPortsLongSyntax(arg map[string]interface{}) (*portsLongSyntax, error) {
	// TODO: Fix spaghetti code
	pls := &portsLongSyntax{}
	v, ok := arg["target"]
	if !ok {
		return nil, fmt.Errorf("Error: map dosen't have target: key")
	}
	tgt, ok := v.(int)
	if !ok {
		return nil, fmt.Errorf("Error: Failed to convert type of value of target: key %v", v)
	}
	if tgt < 0 {
		return nil, fmt.Errorf("Error: Failed to convert value of target: from int to uint because negative %d", tgt)
	}
	pls.target = uint(tgt)

	v, ok = arg["published"]
	if !ok {
		return nil, fmt.Errorf("Error: map dosen't have published: key")
	}
	pub, ok := v.(int)
	if !ok {
		return nil, fmt.Errorf("Error: Failed to convert type of value of published: key %v", v)
	}
	if pub < 0 {
		return nil, fmt.Errorf("Error: Failed to convert value of published: from int to uint because negative %d", pub)
	}
	pls.published = uint(pub)

	v, ok = arg["protocol"]
	if !ok {
		return nil, fmt.Errorf("Error: map dosen't have protocol: key")
	}
	pro, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("Error: Failed to convert type of value of protocol: key %v", v)
	}
	pls.protocol = pro

	v, ok = arg["mode"]
	if !ok {
		return nil, fmt.Errorf("Error: map dosen't have mode: key")
	}
	m, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("Error: Failed to convert type of value of mode: key %v", v)
	}
	pls.mode = m

	return pls, nil
}

func (p *portsLongSyntax) GetSyntaxType() string {
	return p.syntaxType
}

func (p *portsLongSyntax) String() string {
	str := fmt.Sprintf("%starget: %d\n", indents[0], p.target)
	str += fmt.Sprintf("%spublished: %d\n", indents[4], p.published)
	str += fmt.Sprintf("%sprotocol: %s\n", indents[4], p.protocol)
	str += fmt.Sprintf("%smode: %s\n", indents[4], p.mode)
	return str
}

func newPortsShortSyntax(ports string) (*portsShortSyntax, error) {
	pss := &portsShortSyntax{}
	pss.ports = ports

	return pss, nil
}

func (p *portsShortSyntax) GetSyntaxType() string {
	return p.syntaxType
}

func (p *portsShortSyntax) String() string {
	return fmt.Sprintf("\"%s\"\n", p.ports)
}
