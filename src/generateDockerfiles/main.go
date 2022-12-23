package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
)

// main function generates dockerfiles for caddy & the main container
func main() {
	log.Println("Generating dockerfiles...")
	// Load configuration
	config.LoadConfig()
	// Generate dockerfiles
	generate()
	log.Println("Done.")
}

// generate function templates dockerfiles
func generate() {
	templateFile("Dockerfile", "Dockerfile.tmpl")
	templateFile("../caddy-docker/Dockerfile", "caddyDockerFile.tmpl")
}

func templateFile(filename string, templatename string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	template, err := template.ParseFiles(filepath.Join(cwd, "generateDockerfiles/templates/"+templatename))
	if err != nil {
		log.Fatalln(err)
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	template.Execute(f, config.Conf)
	f.Close()
}
