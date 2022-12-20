package app

import (
	"log"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
)

func ConfigurationSanityChecks() {
	if config.Conf.DefaultTLSProvider == "cloudflare" && config.Conf.CloudFlareAPIKey == "" {
		log.Fatal("Error: Default TLS provider set to cloudflare, but no cloudflare API key has been configured.")
	}
}
