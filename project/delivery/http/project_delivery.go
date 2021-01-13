package projecthttpdelivery

import (
	"fmt"
	"net/http"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/gorilla/mux"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ProjectHandler struct {
	PUsecase domain.ProjectUseCase
	// PRepository domain.ProjectRepository
}

func New(r *mux.Router, pu domain.ProjectUseCase) {
	handler := &ProjectHandler{
		PUsecase: pu,
	}

	r.HandleFunc("/project", handler.Fetch()).Methods("GET")
}

func (p *ProjectHandler) Fetch() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		p.PUsecase.FetchAllProjects()
		fmt.Fprintf(w, "fetch project handler")
	}

}
