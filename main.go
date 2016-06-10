package main

import (
	"log"
	"os"
	"runtime"
	
	c "github.com/evanfeenstra/circuitSocket/app/wrappers/config"
	"github.com/evanfeenstra/circuitSocket/app/wrappers/server"
	"github.com/evanfeenstra/circuitSocket/app/route"
	"github.com/evanfeenstra/circuitSocket/app/controller"

)

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Load the configuration file
	c.Load("config"+string(os.PathSeparator)+"config.json", c.Config)

	controller.Configure(c.Config.Client)

	// Start the listener
	server.Run(route.LoadHTTP(), route.LoadHTTPS(), c.Config.Server)
}

