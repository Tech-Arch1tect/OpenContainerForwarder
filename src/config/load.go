package config

import (
	"log"
	"os"
	"strconv"
)

var Conf Config

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

func LoadConfig() {
	var conf Config
	conf.CaddyFileLocation = getDefaultStr("CaddyFileLocation", "/data/config/Caddyfile")
	conf.LoopFrequency = getDefaultInt("LoopFrequency", 15)
	conf.CloudFlareAPIKey = getDefaultStr("CloudFlareAPIKey", "")
	conf.CaddyHost = getDefaultStr("CaddyHost", "http://caddy:2019")
	conf.DefaultTLSProvider = getDefaultStr("DefaultTLSProvider", "default")
	conf.DefaultLogFormat = getDefaultStr("DefaultLogFormat", "console")
	conf.DefaultTrustedProxies = getDefaultStr("DefaultTrustedProxies", "")
	conf.LabelPrefix = "open.container.forwarder"
	conf.WebDashEnabled = getDefaultBool("WebDashEnabled", false)
	conf.CaddyVersion = "2.6.2"
	conf.AlpineVersion = "3.17"
	conf.GoVersion = "1.19.4"
	Conf = conf
}

func getDefaultStr(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

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
