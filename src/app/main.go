package app

/*
TODO:
- Assess whether cmp.equal is appropriate here. Designed for tests, not for production/performance
*/

import (
	"log"
	"time"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/caddyManagement"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/docker"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/extract"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/structs"
)

// RunningContainers stores the extracted information about the running containers
// Exported so that they can be accessed by the web dashboard
var RunningContainers []structs.ContainerExtracts

// GlobalWarnings stores warnings that are generated during the extraction process
// Exported so that they can be accessed by the web dashboard
var GlobalWarnings []string

// Loop is the main loop of the application
// It checks for changes in the running containers and updates the configuration if necessary
func Loop() {
	containers := extract.ExtractInfo(docker.GetContainers(), &GlobalWarnings)
	if !slicesEqual(containers, RunningContainers) {
		// If the extracted container data has changed, generate a new configuration and load it
		log.Println("Container change detected")
		caddyManagement.GenerateConfiguration(containers)
		caddyManagement.LoadConfiguration()
	}
	// Update the running containers variable
	RunningContainers = containers
	// Sleep for the configured amount of time before looping again
	time.Sleep(time.Second * time.Duration(config.Conf.LoopFrequency))
}
