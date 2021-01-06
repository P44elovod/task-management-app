package apiserver

import (
	"log"

	"github.com/P44elovod/task-management-app/config"
)

func Start(config *config.Config) error {
	srv := newServer(config)
	if err := srv.start(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
