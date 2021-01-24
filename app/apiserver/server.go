package apiserver

import (
	"net/http"

	"github.com/P44elovod/task-management-app/config"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *config.Config
	logger *logrus.Logger
	router *mux.Router
}

func (s *Server) newServer(config *config.Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	return http.ListenAndServe(s.config.ServerPort, s.router)
}

func (s *Server) configureLogger() error {

	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}
