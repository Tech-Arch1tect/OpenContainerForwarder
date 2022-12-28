package app

import (
	"testing"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
)

// test CloudflareAPIValidation()
func TestCloudflareAPIValidation(t *testing.T) {
	config.Conf.DefaultTLSProvider = "cloudflare"
	config.Conf.CloudFlareAPIKey = ""
	err := cloudflareAPIValidation()
	if err == nil {
		t.Error("ValidateConfiguration() failed")
	}
}
