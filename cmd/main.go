package main

import (
	"github.com/edgar-lins/gopportunities.git/config"
	"github.com/edgar-lins/gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Failed to initialize configs: %v", err)
		return
	}
	// Initialize the router and start the server
	router.Initialize()
}
