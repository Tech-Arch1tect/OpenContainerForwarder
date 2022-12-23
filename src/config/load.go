package config

import (
	"log"
	"os"
	"strconv"
)

// Conf is the configuration struct exported for use in other packages
var Conf Config

// Config is the configuration struct
type Config struct {
	CaddyFileLocation     string
	LoopFrequency         int
	CloudFlareAPIKey      string
	CaddyHost             string
	DefaultTLSProvider    string
	DefaultLogFormat      string
	DefaultTrustedProxies string
	LabelPrefix           string
	WebDashEnabled        bool
	CaddyVersion          string
	AlpineVersion         string
	GoVersion             string
}

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

// getDefaultStr returns the value of a string environment variable or a default value if the variable is not set
func getDefaultStr(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// getDefaultInt returns the value of a integer environment variable or a default value if the variable is not set
func getDefaultInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal("Failed getting integer value of " + key)
	}
	return v
}

// getDefaultBool returns the value of a boolean environment variable or a default value if the variable is not set
func getDefaultBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	v, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return v
}
