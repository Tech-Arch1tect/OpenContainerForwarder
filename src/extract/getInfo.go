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

func getInfo(rawContainer types.Container, extractedContainers []structs.ContainerStats) (structs.ContainerStats, string, error) {
	isContainerProxied := false
	containerStats := structs.ContainerStats{}
	// set defaults
	containerStats.TLSProvider = config.Conf.DefaultTLSProvider
	containerStats.CloudflareAPIKey = config.Conf.CloudFlareAPIKey
	containerStats.LogFormat = config.Conf.DefaultLogFormat
	containerStats.TrustedProxies = config.Conf.DefaultTrustedProxies
	containerStats.Protocol = "http"
	// get (and overide) values from container labels
	for key, element := range rawContainer.Labels {
		if strings.HasPrefix(key, config.Conf.LabelPrefix) {
			if strings.HasPrefix(key, config.Conf.LabelPrefix+".hostname") {
				isContainerProxied = true
				containerStats.Hostname = append(containerStats.Hostname, element)
			} else if strings.HasPrefix(key, config.Conf.LabelPrefix+".restrictip") {
				if element != "" {
					containerStats.Restrictip = append(containerStats.Restrictip, element)
				}
			} else {
				switch key {
				case config.Conf.LabelPrefix + ".port":
					containerStats.ContainerPort = element
				case config.Conf.LabelPrefix + ".trustedproxies":
					containerStats.TrustedProxies = element
				case config.Conf.LabelPrefix + ".logformat":
					containerStats.LogFormat = element
				case config.Conf.LabelPrefix + ".protocol":
					containerStats.Protocol = element
				case config.Conf.LabelPrefix + ".cloudflareapikey":
					containerStats.CloudflareAPIKey = element
				case config.Conf.LabelPrefix + ".tlsprovider":
					containerStats.TLSProvider = element
				default:
					containerStats.Warnings = append(containerStats.Warnings, "Unrecognised "+config.Conf.LabelPrefix+" label: "+element)
				}
			}

		}
	}
	if isContainerProxied {
		containerStats.HostnameSafe = misc.StripChars(misc.StripChars(containerStats.Hostname[0], ","), " ")
		containerStats.Upstream = docker.GetContainerHostname(rawContainer.ID)
		containerStats.ContainerPort = getPort(&containerStats, rawContainer)
		err := sanityCheckContainer(containerStats, extractedContainers)
		if err != nil {
			return structs.ContainerStats{}, "warning", err
		}
		return containerStats, "", nil
	}
	return structs.ContainerStats{}, "info", errors.New(config.Conf.LabelPrefix + " labels not found")
}
