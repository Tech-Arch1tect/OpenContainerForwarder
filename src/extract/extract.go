package extract

import (
	"github.com/Tech-Arch1tect/OpenContainerForwarder/structs"
	"github.com/docker/docker/api/types"
)

func ExtractInfo(RawContainerData []types.Container, globalWarnings *[]string) []structs.ContainerExtracts {
	extractedContainers := []structs.ContainerExtracts{}
	*globalWarnings = []string{}
	for _, rawContainer := range RawContainerData {
		containerExtracts, elevel, err := getExtract(rawContainer, extractedContainers)
		if err == nil {
			extractedContainers = append(extractedContainers, containerExtracts)
		} else {
			if elevel == "warning" {
				*globalWarnings = append(*globalWarnings, err.Error())
			}
		}
	}
	return extractedContainers
}
