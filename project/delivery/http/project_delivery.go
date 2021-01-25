package projecthttpdelivery

import (
	"encoding/json"
	"net/http"
	"strconv"

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

	r.HandleFunc("/projects", handler.FetchList()).Methods("GET")
	r.HandleFunc("/project/{id:[0-9]+}", handler.GetByID()).Methods("GET")
	r.HandleFunc("/project/{id:[0-9]+}", handler.DelByID()).Methods("DELETE")
	r.HandleFunc("/project/{id:[0-9]+}", handler.UpdateByID()).Methods("PUT")
	r.HandleFunc("/project/new", handler.Create()).Methods("POST")

}

func (p *ProjectHandler) GetByID() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			p.logger.Error(err, id)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid project ID")
			return
		}

		project, err := p.PUsecase.GetByID(uint(id))
		if err != nil {
			p.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Project request went wrong")
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, project)
	}

}

func (p *ProjectHandler) FetchList() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		projectsList, err := p.PUsecase.GetAll()
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

		if err := p.PUsecase.Create(&project); err != nil {
			p.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Requested data is not reached")
			return
		}

		defer r.Body.Close()
		helpers.RespondWithJSON(w, http.StatusCreated, &project)

	}

}

func (p *ProjectHandler) DelByID() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			p.logger.Error(err, id)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid project ID")
			return
		}

		if err := p.PUsecase.DeleteByID(uint(id)); err != nil {
			p.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Project hasn't been deleted")
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, id)
	}

}

func (p *ProjectHandler) UpdateByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			p.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid project ID")
			return
		}

		var project domain.Project
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&project); err != nil {
			p.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		project.ID = uint(id)

		if err := p.PUsecase.UpdateByID(&project); err != nil {
			p.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Project not updated")
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, project)
	}
}
