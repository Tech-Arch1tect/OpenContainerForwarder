package structs

import "github.com/Tech-Arch1tect/OpenContainerForwarder/config"

// ContainerTemplateData is the data passed to the caddyfile template
type ContainerTemplateData struct {
	Containers []ContainerExtracts
	Config     config.Config
}

// ContainerExtracts is the data extracted from a container
type ContainerExtracts struct {
	Hostname         []string `valid:"dns,required"`              // Slice of SNI hostname(s)
	HostnameSafe     string   `valid:"type(string),required"`     // hostname as above, with some chars striped out (used for things like log file name)
	LogFormat        string   `valid:"in(console|json),required"` // caddy log format, console, json etc
	TrustedProxies   string   // space-seperated list of IP (CIDR) addresses to trust as reverse proxies
	Restrictip       []string // slice of ip (CIDR) addresses to restirct access to. If empty, no restriction
	ContainerPort    string   // container listens on this port, traffic forwarded here
	Protocol         string   `valid:"in(http|https),required"` // either https or http - configures what protocol is used to forward traffic to the container
	CloudflareAPIKey string   // If configured globally this will be overridden by the containers labels
	TLSProvider      string   `valid:"in(internal|default|cloudflare),optional"` // How should we optain certs. e.g. cloudflare (dns), direct (default), selfsigned (internal) etc. Set globally and overridden by container-specific label
	Upstream         string   `valid:"required"`                                 // The hostname or IP we are proxying to
	Warnings         []string
}
