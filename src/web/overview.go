package web

import (
	"net/http"
	"strings"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/app"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/caddyManagement"
	"github.com/gin-gonic/gin"
)

func overview(c *gin.Context) {
	TemplateData := getWebTemplatedata()
	c.HTML(http.StatusOK, "index.tmpl", TemplateData)
}

func getWebTemplatedata() webTemplateData {
	tData := webTemplateData{}
	tData.GlobalWarnings = caddyManagement.GlobalWarnings
	for i, container := range app.RunningContainers.Containers {
		containerInfo := webContainerInfo{}
		containerInfo.Hostname = strings.Join(container.Hostname, ",")
		containerInfo.Index = i
		containerInfo.Warnings = container.Warnings
		containerInfo.DisplayLabels = make(map[string]string)
		setLabelIfval(&containerInfo, strings.Join(container.Hostname, ","), "Hostname ("+app.RunningContainers.Config.LabelPrefix+".hostname)")
		setLabelIfval(&containerInfo, container.ContainerPort, "Container Port ("+app.RunningContainers.Config.LabelPrefix+".port)")
		setLabelIfval(&containerInfo, container.LogFormat, "Log Format ("+app.RunningContainers.Config.LabelPrefix+".logformat)")
		setLabelIfval(&containerInfo, strings.Join(container.Restrictip, ","), "IP Lockdown ("+app.RunningContainers.Config.LabelPrefix+".restrictip)")
		setLabelIfval(&containerInfo, container.Protocol, "Reverse Proxy Protocol ("+app.RunningContainers.Config.LabelPrefix+".protocol)")
		setLabelIfval(&containerInfo, container.TLSProvider, "TLS Provider ("+app.RunningContainers.Config.LabelPrefix+".tlsprovider)")
		setLabelIfval(&containerInfo, container.TrustedProxies, "Trusted Proxies ("+app.RunningContainers.Config.LabelPrefix+".trustedproxies)")
		setLabelIfval(&containerInfo, container.Upstream, "Container hostname / ID")
		tData.Containers = append(tData.Containers, containerInfo)
	}
	return tData
}

func setLabelIfval(cInfo *webContainerInfo, value, key string) {
	if value != "" {
		cInfo.DisplayLabels[key] = value
	}
}
