package compose_test

import (
	"testing"

	"github.com/Riki-Okunishi/compose-gen/compose"
)


func TestMain(m *testing.M) {
	m.Run()
}

func TestVersion(t *testing.T) {
	dy := compose.NewDefaultComposeFile()
	if dy.GetVersion() != "3.9" {
		t.Error("Error: not initialized as default version")
	}
	version := "3.5"
	y := compose.NewComposeFile(version)
	if y.GetVersion() != version {
		t.Errorf("Error: not initialized as version %s", version)
	}
}
