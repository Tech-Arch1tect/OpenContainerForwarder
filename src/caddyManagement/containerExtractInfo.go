package caddyManagement

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/docker"
	"github.com/docker/docker/api/types"
)

var Containers []ContainerStats
var GlobalWarnings []string

func ExtractInfo(containers []types.Container) ContainerTemplateData {
	var containerData ContainerTemplateData
	containerData.Config = config.Conf
	Containers = []ContainerStats{}
	GlobalWarnings = []string{}
	for _, container := range containers {
		containerStats, elevel, err := getStats(container)
		if err == nil {
			Containers = append(Containers, containerStats)
		} else {
			if elevel == "warning" {
				GlobalWarnings = append(GlobalWarnings, err.Error())
			}
		}
	}
	containerData.Containers = Containers
	return containerData
}

func getStats(container types.Container) (ContainerStats, string, error) {
	enabled := false
	containerStats := ContainerStats{}
	// set defaults
	containerStats.TLSProvider = config.Conf.DefaultTLSProvider
	containerStats.CloudflareAPIKey = config.Conf.CloudFlareAPIKey
	containerStats.LogFormat = config.Conf.DefaultLogFormat
	containerStats.TrustedProxies = config.Conf.DefaultTrustedProxies
	containerStats.Protocol = "http"
	// get (and overide) values from container labels
	for key, element := range container.Labels {
		if strings.HasPrefix(key, config.Conf.LabelPrefix) {
			if strings.HasPrefix(key, config.Conf.LabelPrefix+".hostname") {
				enabled = true
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
	if enabled {
		containerStats.HostnameSafe = stripChars(stripChars(containerStats.Hostname[0], ","), " ")
		containerStats.Upstream = docker.GetContainerHostname(container.ID)
		containerStats.ContainerPort = getPort(&containerStats, container)
		err := sanityCheckContainer(containerStats)
		if err != nil {
			return ContainerStats{}, "warning", err
		}
		return containerStats, "", nil
	}
	return ContainerStats{}, "info", errors.New(config.Conf.LabelPrefix + " labels not found")
}

func getPort(containerStats *ContainerStats, container types.Container) string {
	// If the container sets the port explicitly, use this
	if containerStats.ContainerPort != "" {
		return containerStats.ContainerPort
	}
	// if there is only only one exposed port, assume this is the desired port
	if len(container.Ports) == 1 {
		return strconv.Itoa(int(container.Ports[0].PrivatePort))
	}
	// if no exposed ports, use port 80 as a best effort and log to inform user
	if len(container.Ports) == 0 {
		containerStats.Warnings = append(containerStats.Warnings, "No exposed ports & "+config.Conf.LabelPrefix+".port not defined. Port 80 is being used as a best effort. Please use the "+config.Conf.LabelPrefix+".port label to define the desired port if not 80.")
		return "80"
	}
	// If there are 2 or more exposed ports, check if any are 80
	webport := 0
	for _, port := range container.Ports {
		if port.PrivatePort == 80 {
			webport = 80
		}
	}
	// if 80 is found, default to this
	if webport == 80 {
		containerStats.Warnings = append(containerStats.Warnings, config.Conf.LabelPrefix+".port not defined. Port 80 was an exposed port so has been used.")
		return "80"
	}
	// If 80 is not found, return the first private port as a best-effort and log a warning that this is likely wrong
	containerStats.Warnings = append(containerStats.Warnings, "Upstream port could not be determined reliably for "+container.Names[0]+". The first exposed port has been found and used. This is likely the wrong port, please use "+config.Conf.LabelPrefix+".port label to configure the correct port.")
	return strconv.Itoa(int(container.Ports[0].PrivatePort))
}
