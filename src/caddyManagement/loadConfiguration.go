package caddyManagement

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
)

func LoadConfiguration() {
	client := &http.Client{}
	data, err := os.Open(config.Conf.CaddyFileLocation)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", config.Conf.CaddyHost+"/load", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "text/caddyfile")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		log.Println("Caddy api not reachable, sleeping for 5 seconds and exiting")
		time.Sleep(time.Second * 5)
		log.Fatal("Exiting")
	}
	if resp.StatusCode == 200 {
		log.Println("Caddy reloaded successfully")
	} else {
		log.Println("Failed to reload caddy. Status code: " + resp.Status)
	}
	defer resp.Body.Close()
}
