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
	"github.com/google/go-cmp/cmp"
)

var RunningContainers []structs.ContainerStats
var GlobalWarnings []string

func Loop() {
	containers := extract.ExtractInfo(docker.GetContainers(), &GlobalWarnings)
	if !slicesEqual(containers, RunningContainers) {
		// containers changed
		log.Println("Container change detected")
		caddyManagement.GenerateConfiguration(containers)
		caddyManagement.LoadConfiguration()
	}
	RunningContainers = containers
	time.Sleep(time.Second * time.Duration(config.Conf.LoopFrequency))
}

func slicesEqual(a, b []structs.ContainerStats) bool {
	return cmp.Equal(a, b)
}
