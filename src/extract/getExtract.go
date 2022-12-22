package extract

import (
	"errors"
	"strings"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/docker"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/misc"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/structs"
	"github.com/docker/docker/api/types"
)

func getExtract(rawContainer types.Container, extractedContainers []structs.ContainerExtracts) (structs.ContainerExtracts, string, error) {
	isContainerProxied := false
	containerExtract := getContainerDefaults()
	// get (and overide) values from container labels
	for key, element := range rawContainer.Labels {
		if strings.HasPrefix(key, config.Conf.LabelPrefix) {
			if strings.HasPrefix(key, config.Conf.LabelPrefix+".hostname") {
				isContainerProxied = true
				containerExtract.Hostname = append(containerExtract.Hostname, element)
			} else if strings.HasPrefix(key, config.Conf.LabelPrefix+".restrictip") {
				if element != "" {
					containerExtract.Restrictip = append(containerExtract.Restrictip, element)
				}
			} else {
				switch key {
				case config.Conf.LabelPrefix + ".port":
					containerExtract.ContainerPort = element
				case config.Conf.LabelPrefix + ".trustedproxies":
					containerExtract.TrustedProxies = element
				case config.Conf.LabelPrefix + ".logformat":
					containerExtract.LogFormat = element
				case config.Conf.LabelPrefix + ".protocol":
					containerExtract.Protocol = element
				case config.Conf.LabelPrefix + ".cloudflareapikey":
					containerExtract.CloudflareAPIKey = element
				case config.Conf.LabelPrefix + ".tlsprovider":
					containerExtract.TLSProvider = element
				default:
					containerExtract.Warnings = append(containerExtract.Warnings, "Unrecognised "+config.Conf.LabelPrefix+" label: "+element)
				}
			}

		}
	}
	if isContainerProxied {
		containerExtract.HostnameSafe = misc.StripChars(misc.StripChars(containerExtract.Hostname[0], ","), " ")
		containerExtract.Upstream = docker.GetContainerHostname(rawContainer.ID)
		containerExtract.ContainerPort = getPort(&containerExtract, rawContainer)
		err := sanityCheckContainer(containerExtract, extractedContainers)
		if err != nil {
			return structs.ContainerExtracts{}, "warning", err
		}
		return containerExtract, "", nil
	}
	return structs.ContainerExtracts{}, "info", errors.New(config.Conf.LabelPrefix + " labels not found")
}
