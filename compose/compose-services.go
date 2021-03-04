package compose

import (
	"fmt"
)

// services is map of service
type serviceEditor interface {
	Build(b string) error
	Image(i string) error
	Volumes(vs ...string) error
	DependsOn(ds ...string) error
}

// service describes one element of services:
type service struct {
	name      string
	version   string
	build     string
	image     string
	volumes   []string
	dependsOn []string
}

func newService(n string, v string) serviceEditor {
	return &service{
		name:    n,
		version: v,
		build: "",
		image: "",
		volumes: []string{},
		dependsOn: []string{},
	}
}

func (s *service) String() string {
	// TODO: define Stringer format
	return ""
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
