package web

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/app"
	"github.com/Tech-Arch1tect/access-log-stats/configuration"
	"github.com/Tech-Arch1tect/access-log-stats/stats"
	"github.com/gin-gonic/gin"
)

func statsOverview(c *gin.Context) {
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "badRequest.tmpl", nil)
		return
	}
	container := app.RunningContainers.Containers[index]
	if container.LogFormat == "" {
		c.HTML(http.StatusBadRequest, "error.tmpl", "log format is empty, no access log found.")
		return
	}
	if container.LogFormat != "console" && container.LogFormat != "json" {
		c.HTML(http.StatusBadRequest, "error.tmpl", "Only console and json log formats are supported for this feature")
		return
	}
	reportData := reportData{}
	reportData.Index = index
	// join all hostnames
	reportData.Hostname = strings.Join(container.Hostname, ", ")
	reportData.Graphs = getGraphs()
	c.HTML(http.StatusOK, "access-log-stats.tmpl", reportData)
}

func statsGetJson(c *gin.Context) {
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "badRequest.tmpl", nil)
		return
	}
	container := app.RunningContainers.Containers[index]
	if container.LogFormat == "" {
		c.HTML(http.StatusBadRequest, "error.tmpl", "log format is empty, no access log found.")
		return
	}
	if container.LogFormat != "console" && container.LogFormat != "json" {
		c.HTML(http.StatusBadRequest, "error.tmpl", "Only console and json log formats are supported for this feature")
		return
	}
	setAccessLogRegexes(container.LogFormat)
	configuration.Config.LogFile = "/var/log/caddy/" + container.HostnameSafe + ".log"
	c.JSON(200, stats.GetStats())
}

func setAccessLogRegexes(logformat string) {
	// Not implemented, will change
	if logformat == "console" {
		configuration.Config.SourceIPRegex = "([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3})"
		configuration.Config.MethodRegex = "\"method\"\\:\\s\"(.*)\"\\,\\s\"host\""
		configuration.Config.URIRegex = "\"uri\"\\:\\s\"(.*)\"\\,\\s\"headers\""
		configuration.Config.ProtocolRegex = "\"proto\"\\:\\s\"(.*)\"\\,\\s\"method\""
		configuration.Config.StatusCodeRegex = "\"status\"\\:\\s([0-9]{3}),\\s"
		configuration.Config.DateRegex = "([0-9]{4}\\/[0-9]{1,2}\\/[0-9]{1,2})"
		configuration.Config.TimeRegex = "\\/[0-9]{2}\\s([0-9]{2}\\:[0-9]{2})"
		configuration.Config.DateFormat = "2006/01/02"
	}
	// not implemented yet, not correct
	if logformat == "json" {
		configuration.Config.SourceIPRegex = "([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3})"
		configuration.Config.MethodRegex = "\"method\"\\:\\s\"(.*)\"\\,\\s\"host\""
		configuration.Config.URIRegex = "\"uri\"\\:\\s\"(.*)\"\\,\\s\"headers\""
		configuration.Config.ProtocolRegex = "\"proto\"\\:\\s\"(.*)\"\\,\\s\"method\""
		configuration.Config.StatusCodeRegex = "\"status\"\\:\\s([0-9]{3}),\\s"
		configuration.Config.DateRegex = "([0-9]{4}\\/[0-9]{1,2}\\/[0-9]{1,2})"
		configuration.Config.TimeRegex = "\\/[0-9]{2}\\s([0-9]{2}\\:[0-9]{2})"
		configuration.Config.DateFormat = "2006/01/02"
	}
}

func getGraphs() map[string]map[string]string {
	mapData := make(map[string]map[string]string)
	mapData = getMapData("SourceIPs", "10", "bar", mapData)
	mapData = getMapData("Methods", "0", "bar", mapData)
	mapData = getMapData("Protocols", "10", "bar", mapData)
	mapData = getMapData("URI", "10", "bar", mapData)
	mapData = getMapData("StatusCodes", "0", "bar", mapData)
	mapData = getMapData("Date", "0", "line", mapData)
	mapData = getMapData("ByHour", "0", "line", mapData)
	return mapData
}

func getMapData(name, limit string, graphtype string, mapData map[string]map[string]string) map[string]map[string]string {
	mapData[name] = make(map[string]string)
	mapData[name]["Name"] = name
	mapData[name]["Limit"] = limit
	mapData[name]["GraphType"] = graphtype
	return mapData
}
