package caddyManagement

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/structs"
)

// GenerateConfiguration generates the caddyfile from the extracted container data
func GenerateConfiguration(containers []structs.ContainerExtracts) {
	tData := structs.ContainerTemplateData{}
	tData.Containers = containers
	tData.Config = config.Conf
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	t, err := template.New("caddyfile.tmpl").Funcs(template.FuncMap{"joinStrings": joinStrings}).ParseFiles(filepath.Join(cwd, "templates/caddy/caddyfile.tmpl"))
	if err != nil {
		log.Fatalln(err)
	}
	f, err := os.Create(config.Conf.CaddyFileLocation)
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(f, tData)
	if err != nil {
		log.Fatalln(err)
	}
	f.Close()
}

// LoadConfiguration loads the caddyfile into caddy using the caddy api
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
