package taskhttpdelivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type TaskHandler struct {
	TUsecase domain.TaskUseCase
	logger   *logrus.Logger
}

func New(r *mux.Router, log *logrus.Logger, tu domain.TaskUseCase) {
	handler := &TaskHandler{
		TUsecase: tu,
		logger:   log,
	}

	r.HandleFunc("/task/new", handler.Create()).Methods("POST")
	r.HandleFunc("/task/{id:[0-9]+}", handler.GetByID()).Methods("GET")
	r.HandleFunc("/comment/{id:[0-9]+}", handler.DeleteByID()).Methods("DELETE")
	r.HandleFunc("/comment/{id:[0-9]+}", handler.UpdateByID()).Methods("PUT")

}

func (th *TaskHandler) Create() http.HandlerFunc {
	var task domain.Task
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&task); err != nil {
			th.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if err := th.TUsecase.CreateTask(&task); err != nil {
			th.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Task not created")
			return
		}

		defer r.Body.Close()

		helpers.RespondWithJSON(w, http.StatusCreated, &task)

	}
}

func (th *TaskHandler) GetByID() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			th.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid task ID")
			return
		}

		projectsList, err := th.TUsecase.GetTaskWithCommentByID(uint(id))
		if err != nil {
			th.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Task request went wrong")
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, projectsList)
	}
}

func (th *TaskHandler) UpdateByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			th.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid task ID")
			return
		}

		var task domain.Task
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&task); err != nil {
			th.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		task.ID = uint(id)

		if err := th.TUsecase.Update(&task); err != nil {
			th.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, task)
	}
}

func (th *TaskHandler) DeleteByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			th.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid task ID")
			return
		}

		if err := th.TUsecase.DeleteByID(uint(id)); err != nil {
			th.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Task hasn't been deleted")
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, id)

	}
}
