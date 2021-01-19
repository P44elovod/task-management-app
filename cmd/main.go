package main

import (
	"flag"
	"log"

	"github.com/P44elovod/task-management-app/app/apiserver"
	"github.com/P44elovod/task-management-app/config"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./", "config file path")
}

func main() {
	a := apiserver.Api{}

	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	a.Start(config)
}
