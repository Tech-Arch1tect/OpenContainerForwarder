package app

import (
	"errors"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
)

// ValidateConfiguration checks for errors in the configuration
func ValidateConfiguration() {
	err := cloudflareAPIValidation()
	if err != nil {
		panic(err)
	}
}

// cloudflareAPIValidation checks if the cloudflare api key is set if the default tls provider is set to cloudflare
func cloudflareAPIValidation() error {
	if config.Conf.DefaultTLSProvider == "cloudflare" && config.Conf.CloudFlareAPIKey == "" {
		return errors.New("default tls provider set to cloudflare, but no cloudflare api key has been configured")
	}
	return nil
}
