package extract

import (
	"strconv"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/structs"
	"github.com/docker/docker/api/types"
)

func getPort(containerExtract *structs.ContainerExtracts, container types.Container) string {
	// If the container sets the port explicitly, use this
	if containerExtract.ContainerPort != "" {
		return containerExtract.ContainerPort
	}
	// if there is only only one exposed port, assume this is the desired port
	if len(container.Ports) == 1 {
		return strconv.Itoa(int(container.Ports[0].PrivatePort))
	}
	// if no exposed ports, use port 80 as a best effort and log to inform user
	if len(container.Ports) == 0 {
		containerExtract.Warnings = append(containerExtract.Warnings, "No exposed ports & "+config.Conf.LabelPrefix+".port not defined. Port 80 is being used as a best effort. Please use the "+config.Conf.LabelPrefix+".port label to define the desired port if not 80.")
		return "80"
	}
	// If there are 2 or more exposed ports, check if any are 80
	foundPort := 0
	for _, port := range container.Ports {
		if port.PrivatePort == 80 {
			foundPort = 80
		}
	}
	// if 80 is found, default to this
	if foundPort == 80 {
		containerExtract.Warnings = append(containerExtract.Warnings, config.Conf.LabelPrefix+".port not defined. Port 80 was an exposed port so has been used.")
		return "80"
	}
	// If 80 is not found, return the first private port as a best-effort and log a warning that this is likely wrong
	containerExtract.Warnings = append(containerExtract.Warnings, "Upstream port could not be determined reliably for "+container.Names[0]+". The first exposed port has been found and used. This is likely the wrong port, please use "+config.Conf.LabelPrefix+".port label to configure the correct port.")
	return strconv.Itoa(int(container.Ports[0].PrivatePort))
}
