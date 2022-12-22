package extract

import (
	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/structs"
)

func getContainerDefaults() structs.ContainerExtracts {
	containerExtract := structs.ContainerExtracts{}
	// set defaults
	containerExtract.TLSProvider = config.Conf.DefaultTLSProvider
	containerExtract.CloudflareAPIKey = config.Conf.CloudFlareAPIKey
	containerExtract.LogFormat = config.Conf.DefaultLogFormat
	containerExtract.TrustedProxies = config.Conf.DefaultTrustedProxies
	containerExtract.Protocol = "http"
	return containerExtract
}
