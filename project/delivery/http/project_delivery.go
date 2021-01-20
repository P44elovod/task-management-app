package projecthttpdelivery

import (
	"encoding/json"
	"net/http"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
	"github.com/gorilla/mux"
)

type ProjectHandler struct {
	PUsecase domain.ProjectUsecase
	// PRepository domain.ProjectRepository
}

func New(r *mux.Router, pu domain.ProjectUsecase) {
	handler := &ProjectHandler{
		PUsecase: pu,
	}

	r.HandleFunc("/projects", handler.Fetch()).Methods("GET")
	r.HandleFunc("/project/new", handler.Create()).Methods("POST")

}

func (p *ProjectHandler) Fetch() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		projectsList, err := p.PUsecase.FetchAllProjects()
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Project list request went wrong")
		}

		helpers.RespondWithJSON(w, http.StatusOK, projectsList)
	}

}

func (p *ProjectHandler) Create() http.HandlerFunc {
	var project domain.Project

	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&project); err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		}
		err := p.PUsecase.CreateProject(&project)
		if err != nil {
			helpers.FailOnError(err, "WTF")
			helpers.RespondWithError(w, http.StatusInternalServerError, "Rquested data is not reached")
		}

		defer r.Body.Close()
		helpers.RespondWithJSON(w, http.StatusCreated, &project)

	}

}
