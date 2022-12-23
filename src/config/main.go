package config

import "github.com/Tech-Arch1tect/OpenContainerForwarder/structs"

// Conf is the configuration struct exported for use in other packages
var Conf structs.Config

// LoadConfig loads the configuration from environment variables
func LoadConfig() {
	Conf.CaddyFileLocation = getDefaultStr("CaddyFileLocation", "/data/config/Caddyfile")
	Conf.LoopFrequency = getDefaultInt("LoopFrequency", 15)
	Conf.CloudFlareAPIKey = getDefaultStr("CloudFlareAPIKey", "")
	Conf.CaddyHost = getDefaultStr("CaddyHost", "http://caddy:2019")
	Conf.DefaultTLSProvider = getDefaultStr("DefaultTLSProvider", "default")
	Conf.DefaultLogFormat = getDefaultStr("DefaultLogFormat", "console")
	Conf.DefaultTrustedProxies = getDefaultStr("DefaultTrustedProxies", "")
	Conf.LabelPrefix = "open.container.forwarder"
	Conf.WebDashEnabled = getDefaultBool("WebDashEnabled", false)
	Conf.CaddyVersion = "2.6.2"
	Conf.AlpineVersion = "3.17"
	Conf.GoVersion = "1.19.4"
}
