package extract

import (
	"testing"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/structs"
	"github.com/docker/docker/api/types"
)

// test getExtract()
func TestGetExtract(t *testing.T) {
	// load config
	config.LoadConfig()
	// create a raw container
	rawContainer := types.Container{}
	rawContainer.ID = "test"
	rawContainer.Labels = map[string]string{
		"open.container.forwarder.hostname": "test.local",
		"open.container.forwarder.port":     "80",
	}
	ContainerExtracts := []structs.ContainerExtracts{}
	Container, _, err := getExtract(rawContainer, ContainerExtracts)
	// Fail on an error
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	// Fail if hostname is not test.local
	if Container.Hostname[0] != "test.local" {
		t.Errorf("Hostname: %v", Container.Hostname)
	}
	// Fail if port is not 80
	if Container.ContainerPort != "80" {
		t.Errorf("Port: %v", Container.ContainerPort)
	}
}

// test getExtract() when no hostname is present
func TestGetExtract_hostname(t *testing.T) {
	// load config
	config.LoadConfig()
	// create a raw container
	rawContainer := types.Container{}
	rawContainer.ID = "test"
	rawContainer.Labels = map[string]string{
		"open.container.forwarder.port": "80",
	}
	ContainerExtracts := []structs.ContainerExtracts{}
	_, _, err := getExtract(rawContainer, ContainerExtracts)
	// Fail if no error is returned
	if err == nil {
		t.Errorf("Error: %v", err)
	}
}

// test getExtract() when no port is present
func TestGetExtract_port(t *testing.T) {
	// load config
	config.LoadConfig()
	// create a raw container
	rawContainer := types.Container{}
	rawContainer.ID = "test"
	rawContainer.Labels = map[string]string{
		"open.container.forwarder.hostname": "test.local",
	}
	ContainerExtracts := []structs.ContainerExtracts{}
	container, _, err := getExtract(rawContainer, ContainerExtracts)
	// Fail if there is an error
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	// Fail if port is not 80
	if container.ContainerPort != "80" {
		t.Errorf("Port: %v", ContainerExtracts[0].ContainerPort)
	}
}

// test getExtract() adds a warning when unreconised labels are present
func TestGetExtract_unrecognised(t *testing.T) {
	// load config
	config.LoadConfig()
	// create a raw container
	rawContainer := types.Container{}
	rawContainer.ID = "test"
	rawContainer.Labels = map[string]string{
		"open.container.forwarder.hostname": "test.local",
		"open.container.forwarder.port":     "80",
		"open.container.forwarder.test":     "test",
	}
	ContainerExtracts := []structs.ContainerExtracts{}
	Container, _, err := getExtract(rawContainer, ContainerExtracts)
	// Fail if there is an error
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	// Fail if there are no warnings
	if len(Container.Warnings) == 0 {
		t.Errorf("Warnings: %v", Container.Warnings)
	}
}

// test getPort()
func TestGetPort(t *testing.T) {
	// load config
	config.LoadConfig()
	// create a raw container & container extracts
	rawContainer := types.Container{}
	containerExtracts := structs.ContainerExtracts{}
	// get the port
	containerExtracts.ContainerPort = "8080"
	port := getPort(&containerExtracts, rawContainer)
	// Fail if port is not 8080
	if port != "8080" {
		t.Errorf("Port: %v", port)
	}
	// test getPort() when no port is present
	containerExtracts.ContainerPort = ""
	port = getPort(&containerExtracts, rawContainer)
	// Fail if port is not 80
	if port != "80" {
		t.Errorf("Port: %v", port)
	}
}
