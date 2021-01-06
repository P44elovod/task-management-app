package apiserver

import (
	"fmt"
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

func newServer(config *config.Config) *Server {
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

	s.configureRouter()

	s.logger.Info("start api server")
	return http.ListenAndServe(s.config.ServerPort, s.router)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.handleHello())

}

func (s *Server) configureLogger() error {

	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Server) handleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "hello")
		s.logger.Info("all good")
	}
}
