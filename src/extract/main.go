package extract

import (
	"github.com/Tech-Arch1tect/OpenContainerForwarder/structs"
	"github.com/docker/docker/api/types"
)

// ExtractInfo extracts information from a struct of docker containers and returns a struct of extracted container data
func ExtractInfo(RawContainerData []types.Container, globalWarnings *[]string) []structs.ContainerExtracts {
	// Init variable to hold extracted container data
	extractedContainers := []structs.ContainerExtracts{}
	// Init variable to hold warnings from the extract process ( which can not be related to a specific container)
	*globalWarnings = []string{}

	// Loop through all containers and extract data if labels are present
	for _, rawContainer := range RawContainerData {
		containerExtracts, elevel, err := getExtract(rawContainer, extractedContainers)
		// store extracted container data if no errors
		if err == nil {
			extractedContainers = append(extractedContainers, containerExtracts)
		} else {
			// warning level errors are not fatal, but should be stored for display
			if elevel == "warning" {
				*globalWarnings = append(*globalWarnings, err.Error())
			}
		}
	}
	// return extracted container data
	return extractedContainers
}
