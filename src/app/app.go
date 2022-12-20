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
	"github.com/google/go-cmp/cmp"
)

var RunningContainers caddyManagement.ContainerTemplateData

func Loop() {
	containers := caddyManagement.Collect()
	if !slicesEqual(containers.Containers, RunningContainers.Containers) {
		// containers changed
		log.Println("Container change detected")
		caddyManagement.GenerateConfiguration(containers)
		caddyManagement.LoadConfiguration()
	}
	RunningContainers = containers
	time.Sleep(time.Second * time.Duration(config.Conf.LoopFrequency))
}

func slicesEqual(a, b []caddyManagement.ContainerStats) bool {
	return cmp.Equal(a, b)
}
