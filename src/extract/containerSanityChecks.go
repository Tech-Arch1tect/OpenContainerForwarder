package extract

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/structs"
	"github.com/asaskevich/govalidator"
)

func sanityCheckContainer(newContainer structs.ContainerStats, containers []structs.ContainerStats) error {
	_, err := govalidator.ValidateStruct(newContainer)
	if err != nil {
		return err
	}
	// check for duplicate hostnames
	for _, container := range containers {
		for _, hostname := range container.Hostname {
			for _, newHostname := range newContainer.Hostname {
				if hostname == newHostname {
					return errors.New("Duplicate hostname found: " + hostname)
				}
			}
		}
	}
	// check restrictip is a valid CIDR
	for _, restrictip := range newContainer.Restrictip {
		if !govalidator.IsCIDR(restrictip) {
			return errors.New("Invalid restrictip: " + restrictip)
		}
	}
	// check port is an int stored as a string
	port, err := strconv.Atoi(newContainer.ContainerPort)
	if err != nil {
		return errors.New("Invalid port: " + newContainer.ContainerPort)
	}
	// check port is not 0
	if port == 0 {
		return errors.New("Invalid port: " + newContainer.ContainerPort)
	}
	// Break trusted proxies into a slice by space character and check each is valid cidr
	if newContainer.TrustedProxies != "" {
		for _, trustedProxy := range strings.Split(newContainer.TrustedProxies, " ") {
			if !govalidator.IsCIDR(trustedProxy) {
				return errors.New("Invalid trusted proxy: " + trustedProxy)
			}
		}
	}

	return nil
}
