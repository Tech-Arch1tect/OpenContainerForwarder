package extract

import (
	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/structs"
)

// getContainerDefaults returns a structs.ContainerExtracts with default container values
func getContainerDefaults() structs.ContainerExtracts {
	// Init empty variable to hold extracted container data
	containerExtract := structs.ContainerExtracts{}
	// set defaults
	containerExtract.TLSProvider = config.Conf.DefaultTLSProvider
	containerExtract.CloudflareAPIKey = config.Conf.CloudFlareAPIKey
	containerExtract.LogFormat = config.Conf.DefaultLogFormat
	containerExtract.TrustedProxies = config.Conf.DefaultTrustedProxies
	containerExtract.Protocol = "http"
	// return default container data
	return containerExtract
}
