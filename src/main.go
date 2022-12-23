package main

import (
	"github.com/Tech-Arch1tect/OpenContainerForwarder/app"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/config"
	"github.com/Tech-Arch1tect/OpenContainerForwarder/web"
)

func main() {
	// Load configuration
	config.LoadConfig()
	// Check for errors in configuration
	app.ValidateConfiguration()
	// Start web dashboard
	web.StartWeb()
	// Start main loop
	for {
		app.Loop()
	}
}
