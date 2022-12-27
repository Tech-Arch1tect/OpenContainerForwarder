package extract

import (
	"errors"
	"sort"
	"strings"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/docker"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/misc"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/structs"
	"github.com/docker/docker/api/types"
)

// getExtract extracts information from a single docker container and returns a struct of extracted container data if labels are present
func getExtract(rawContainer types.Container, extractedContainers []structs.ContainerExtracts) (structs.ContainerExtracts, string, error) {
	// isContainerProxied is a flag to indicate if the container contains any labels that are relevant to the proxy
	isContainerProxied := false
	// init containerExtracts struct with default values
	containerExtract := getContainerDefaults()
	// get (and override) values from container labels
	for key, element := range rawContainer.Labels {
		// check if label is relevant to the proxy
		if strings.HasPrefix(key, config.Conf.LabelPrefix) {
			// check if label is a hostname label (including multiple hostnames)
			// check if label is a restrictip label (including multiple restrictips)
			// else check if label is a single value label
			if strings.HasPrefix(key, config.Conf.LabelPrefix+".hostname") {
				isContainerProxied = true
				containerExtract.Hostname = append(containerExtract.Hostname, element)
			} else if strings.HasPrefix(key, config.Conf.LabelPrefix+".restrictip") {
				if element != "" {
					containerExtract.Restrictip = append(containerExtract.Restrictip, element)
				}
			} else {
				// check if label is a single value label
				// validation of values is performed in the sanityCheckContainer function later
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
					// if label is not recognised, add it to the warnings slice
					containerExtract.Warnings = append(containerExtract.Warnings, "Unrecognised "+config.Conf.LabelPrefix+" label: "+element)
				}
			}

		}
	}
	// if container should be proxied, perform sanity checks and return containerExtracts struct
	if isContainerProxied {
		// Sort slices to prevent detecting container changes when order of slices changes
		sort.Strings(containerExtract.Hostname)
		sort.Strings(containerExtract.Restrictip)
		sort.Strings(containerExtract.Warnings)

		containerExtract.HostnameSafe = misc.StripChars(misc.StripChars(containerExtract.Hostname[0], ","), " ")
		containerExtract.Upstream = docker.GetContainerHostname(rawContainer.ID)
		containerExtract.ContainerPort = getPort(&containerExtract, rawContainer)
		// perform validation of values
		err := validateContainer(containerExtract, extractedContainers)
		if err != nil {
			// if sanity check fails, return error and empty containerExtracts struct
			return structs.ContainerExtracts{}, "warning", err
		}
		// if sanity check passes, return containerExtracts struct without error
		return containerExtract, "", nil
	}
	// if container should not be proxied, return empty containerExtracts struct with info error level so that it is not stored in globalWarnings
	return structs.ContainerExtracts{}, "info", errors.New(config.Conf.LabelPrefix + " labels not found")
}
