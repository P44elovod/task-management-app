package taskhttpdelivery

import (
	"encoding/json"
	"net/http"

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

		err := th.TUsecase.CreateTask(&task)
		if err != nil {
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
		projectsList, err := th.TUsecase.GetTaskWithCommentByID(vars["id"])
		if err != nil {
			th.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Task request went wrong")
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, projectsList)
	}

}
