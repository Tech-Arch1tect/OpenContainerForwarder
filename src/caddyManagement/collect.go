package caddyManagement

import "github.com/Tech-Arch1tect/OpenContainerForwarder/docker"

func Collect() ContainerTemplateData {
	containers := docker.GetContainers()
	return ExtractInfo(containers)
}
