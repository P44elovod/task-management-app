package taskhttpdelivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
	"github.com/gorilla/mux"
)

type TaskHandler struct {
	TUsecase domain.TaskUseCase
}

func New(r *mux.Router, tu domain.TaskUseCase) {
	handler := &TaskHandler{
		TUsecase: tu,
	}

	r.HandleFunc("/task/new", handler.Create()).Methods("POST")
	r.HandleFunc("/task/{id:[0-9]+}", handler.GetByID()).Methods("GET")

}

func (th *TaskHandler) Create() http.HandlerFunc {
	var task domain.Task
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&task); err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		}

		err := th.TUsecase.CreateTask(&task)
		if err != nil {
			fmt.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Task not created")
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
			helpers.RespondWithError(w, http.StatusBadRequest, "Task request went wrong")
		}

		helpers.RespondWithJSON(w, http.StatusOK, projectsList)
	}

}
