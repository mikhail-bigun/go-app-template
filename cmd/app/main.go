package main

import (
	"goapptemplate/config"
	"goapptemplate/internal/app"

	"log"

	"github.com/spf13/pflag"
)

//	@title			Books API
//	@version		0.1.0
//	@description	Go app template books API.

//	@contact.name	John Doe
//	@contact.email	johndoe@cia.gov

//	@schemes	http https
func main() {
	// ________________________________________________________________________
	// Parse cli args to config
	filepath := pflag.StringP("config", "c", "", "configuration filepath (default: None)")
	pflag.Parse()
	// ________________________________________________________________________
	// Load config
	cfg, err := config.NewAppCfg(*filepath)
	if err != nil {
		log.Fatalf("cannot load config: %s", err)
	}
	// ________________________________________________________________________
	// Run app
	app.Run(cfg)
}
