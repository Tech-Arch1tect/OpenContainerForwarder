package extract

import (
	"github.com/Tech-Arch1tect/OpenContainerForwarder/structs"
	"github.com/docker/docker/api/types"
)

func ExtractInfo(RawContainerData []types.Container, globalWarnings *[]string) []structs.ContainerStats {
	extractedContainers := []structs.ContainerStats{}
	*globalWarnings = []string{}
	for _, rawContainer := range RawContainerData {
		containerStats, elevel, err := getInfo(rawContainer, extractedContainers)
		if err == nil {
			extractedContainers = append(extractedContainers, containerStats)
		} else {
			if elevel == "warning" {
				*globalWarnings = append(*globalWarnings, err.Error())
			}
		}
	}
	return extractedContainers
}
