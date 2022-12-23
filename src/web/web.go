package web

/*
This is extremely early / not close to being properly functional
TODO:
- Add stats from access log
- Everything else
*/

import (
	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
	"github.com/gin-gonic/gin"
)

// StartWeb starts the web dashboard if enabled
func StartWeb() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/web/*")
	// routes
	r.GET("/", overview)
	r.GET("/stats/:index", statsOverview)
	r.GET("/stats/:index/json", statsGetJson)
	r.GET("/debug/dump/heap", debugHeap)
	r.GET("/debug/dump/containers", debugContainers)
	// run gin via goroutine
	if config.Conf.WebDashEnabled {
		go r.Run()
	}

}
