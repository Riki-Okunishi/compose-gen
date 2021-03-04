package compose_test

import (
	"testing"

	"github.com/Riki-Okunishi/compose-gen/compose"
)


func TestServicesBuild(t *testing.T) {
	yml := compose.NewDefaultComposeFile()

	yml.Services("app").Build("samplePath")
	yml.Services("app").Build("path")
}

func TestServicesDependsOn(t *testing.T) {
	yml := compose.NewDefaultComposeFile()

	yml.Services("app").DependsOn("db", "web")
}