package caddyManagement

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
)

func GenerateConfiguration(templateData ContainerTemplateData) {
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
	err = t.Execute(f, templateData)
	if err != nil {
		log.Fatalln(err)
	}
	f.Close()
}

func joinStrings(s []string, sep string) string {
	return strings.Join(s, sep)
}
