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
	r.HandleFunc("/task/{id:[0-9]+}", handler.DeleteByID()).Methods("DELETE")
	r.HandleFunc("/task/{id:[0-9]+}", handler.UpdateByID()).Methods("PUT")
	r.HandleFunc("/task/priority", handler.UpdatePriority()).Methods("PUT")

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

		taskList, err := th.TUsecase.GetTaskWithCommentByID(uint(id))
		if err != nil {
			th.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Task request went wrong")
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, taskList)
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

		response := map[string]interface{}{
			"id":     id,
			"status": "success",
		}

		helpers.RespondWithJSON(w, http.StatusOK, response)

	}
}

func (th *TaskHandler) UpdatePriority() http.HandlerFunc {

	type Priority struct {
		ID       uint `json:"id"`
		Priority uint `json:"priority"`
	}

	type Priorities struct {
		Priorities []Priority `json:"priorities"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		var priorityList Priorities
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&priorityList); err != nil {
			th.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()

		priorityMap := make(map[uint]uint)
		for i := 0; i < len(priorityList.Priorities); i++ {
			priorityMap[priorityList.Priorities[i].ID] = priorityList.Priorities[i].Priority
		}

		if err := th.TUsecase.UpdateUpdatePriority(priorityMap); err != nil {
			th.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Columns postions not updated")
			return
		}

		response := map[string]interface{}{
			"priorities": priorityMap,
			"status":     "success",
		}

		helpers.RespondWithJSON(w, http.StatusOK, response)
	}
}
