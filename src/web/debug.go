package web

import (
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"time"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/app"
	"github.com/gin-gonic/gin"
)

func debugHeap(c *gin.Context) {
	f, err := os.Create("/dumps/" + time.Now().String() + ".heap")
	if err != nil {
		log.Fatalf("%v", err)
	}
	pprof.WriteHeapProfile(f)
	f.Close()
	c.JSON(http.StatusOK, "Heap dumped")
}

func debugContainers(c *gin.Context) {
	c.JSON(http.StatusOK, app.RunningContainers)
}
