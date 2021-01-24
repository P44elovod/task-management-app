package projecthttpdelivery

import (
	"encoding/json"
	"net/http"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ProjectHandler struct {
	PUsecase domain.ProjectUsecase
	logger   *logrus.Logger
}

func New(r *mux.Router, log *logrus.Logger, pu domain.ProjectUsecase) {
	handler := &ProjectHandler{
		PUsecase: pu,
	}

	r.HandleFunc("/projects", handler.Fetch()).Methods("GET")
	r.HandleFunc("/project/{id:[0-9]+}", handler.GetByID()).Methods("GET")
	r.HandleFunc("/project/new", handler.Create()).Methods("POST")

}

func (p *ProjectHandler) GetByID() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		project, err := p.PUsecase.GetProjectByID(vars["id"])
		if err != nil {
			p.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Project request went wrong")
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, project)
	}

}

func (p *ProjectHandler) Fetch() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		projectsList, err := p.PUsecase.FetchAllProjects()
		if err != nil {
			p.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Project list request went wrong")
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, projectsList)
	}

}

func (p *ProjectHandler) Create() http.HandlerFunc {
	var project domain.Project

	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&project); err != nil {
			p.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		err := p.PUsecase.CreateProject(&project)
		if err != nil {
			p.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Rquested data is not reached")
			return
		}

		defer r.Body.Close()
		helpers.RespondWithJSON(w, http.StatusCreated, &project)

	}

}
