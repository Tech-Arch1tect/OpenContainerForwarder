package web

import (
	"net/http"
	"strings"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/app"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
	"github.com/gin-gonic/gin"
)

// overview is the main page of the web dashboard
func overview(c *gin.Context) {
	TemplateData := getWebTemplatedata()
	c.HTML(http.StatusOK, "index.tmpl", TemplateData)
}

// getWebTemplatedata returns the data for the web dashboard
func getWebTemplatedata() webTemplateData {
	tData := webTemplateData{}
	tData.GlobalWarnings = app.GlobalWarnings
	for i, container := range app.RunningContainers {
		containerInfo := webContainerInfo{}
		containerInfo.Hostname = strings.Join(container.Hostname, ",")
		containerInfo.Index = i
		containerInfo.Warnings = container.Warnings
		containerInfo.DisplayLabels = make(map[string]string)
		setLabelIfval(&containerInfo, strings.Join(container.Hostname, ","), "Hostname ("+config.Conf.LabelPrefix+".hostname)")
		setLabelIfval(&containerInfo, container.ContainerPort, "Container Port ("+config.Conf.LabelPrefix+".port)")
		setLabelIfval(&containerInfo, container.LogFormat, "Log Format ("+config.Conf.LabelPrefix+".logformat)")
		setLabelIfval(&containerInfo, strings.Join(container.Restrictip, ","), "IP Lockdown ("+config.Conf.LabelPrefix+".restrictip)")
		setLabelIfval(&containerInfo, container.Protocol, "Reverse Proxy Protocol ("+config.Conf.LabelPrefix+".protocol)")
		setLabelIfval(&containerInfo, container.TLSProvider, "TLS Provider ("+config.Conf.LabelPrefix+".tlsprovider)")
		setLabelIfval(&containerInfo, container.TrustedProxies, "Trusted Proxies ("+config.Conf.LabelPrefix+".trustedproxies)")
		setLabelIfval(&containerInfo, container.Upstream, "Container hostname / ID")
		tData.Containers = append(tData.Containers, containerInfo)
	}
	return tData
}

// setLabelIfval sets a label in the webContainerInfo struct if the value is not empty
func setLabelIfval(cInfo *webContainerInfo, value, key string) {
	if value != "" {
		cInfo.DisplayLabels[key] = value
	}
}
