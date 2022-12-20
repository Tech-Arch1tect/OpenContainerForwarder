package main

import (
	"github.com/Tech-Arch1tect/OpenContainerForwarder/app"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/web"
)

func main() {
	config.LoadConfig()
	app.ConfigurationSanityChecks()
	web.StartWeb()
	for {
		app.Loop()
	}
}
