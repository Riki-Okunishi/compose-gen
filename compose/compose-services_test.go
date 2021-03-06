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

func TestServicesPortsValidation(t *testing.T) {
	yml := compose.NewDefaultComposeFile()

	// see https://docs.docker.com/compose/compose-file/compose-file-v3/#ports
	correctPorts := []string{
		"3000",
		"3000-3005",
		"8000:8000",
		"9090-9091:8080-8081",
		"49100:22",
		"127.0.0.1:8001:8001",
		"127.0.0.0:5000-5010:5000-5010",
		"127.0.0.1::5000",
		"6060:6060/udp",
		"12400-12500:1240",
	}

	for i, ports := range correctPorts {
		err := yml.Services("app").Ports(ports)
		if err != nil {
			t.Errorf("Error in test case[%d]: Failed to add a correct ports: value as short syntax %s", i, ports)
		}
	}
}

func TestServicesPortsArgumentsType(t *testing.T) {
	yml := compose.NewDefaultComposeFile()
	app := yml.Services("app")

	err := app.Ports("8080:80")
	if err != nil {
		t.Errorf("Error: Failed to add a correct ports: value as short syntax 8080:80")
	}

	err = app.Ports("8080:80", "8081:81")
	if err != nil {
		t.Errorf("Error: Failed to add a correct ports: value as short syntax 8080:80, 8081:81")
	}

	portList := []string{
		"8080:80",
		"8081:81",
	}
	err = app.Ports(portList)
	if err != nil {
		t.Errorf("Error: Failed to add a correct ports: value as short syntax %v", portList)
	}

	pls := map[string]interface{}{
		"target":    80,
		"published": 8080,
		"protocol":  "tcp",
		"mode":      "host",
	}
	err = app.Ports(pls)
	if err != nil {
		t.Errorf("Error: Failed to add a correct ports: value as long syntax %v\n detail: %s", pls, err)
	}

	plsList := []map[string]interface{}{
		map[string]interface{}{
			"target":    80,
			"published": 8080,
			"protocol":  "tcp",
			"mode":      "host",
		},
		map[string]interface{}{
			"target":    81,
			"published": 8081,
			"protocol":  "udp",
			"mode":      "host",
		},
	}
	err = app.Ports(plsList)
	if err != nil {
		t.Errorf("Error: Failed to add a correct ports: value as long syntax %v", plsList)
	}
}
