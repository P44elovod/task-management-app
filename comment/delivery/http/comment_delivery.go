package commenthttpdelivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type CommentHandler struct {
	CMUsecase domain.CommentUseCase
	CMRepo    domain.CommentRepository
	logger    *logrus.Logger
}

func New(r *mux.Router, log *logrus.Logger, cmu domain.CommentUseCase, cmr domain.CommentRepository) {
	handler := &CommentHandler{
		CMUsecase: cmu,
		CMRepo:    cmr,
		logger:    log,
	}

	r.HandleFunc("/comment/new", handler.Create()).Methods("POST")
	r.HandleFunc("/comment/{id:[0-9]+}", handler.DeleteByID()).Methods("DELETE")
	r.HandleFunc("/comment/{id:[0-9]+}", handler.UpdateByID()).Methods("PUT")

}

func (ch *CommentHandler) Create() http.HandlerFunc {
	var comment domain.Comment

	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&comment); err != nil {
			ch.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		err := ch.CMRepo.StoreComment(&comment)
		if err != nil {
			ch.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Comment not created")
			return
		}
		defer r.Body.Close()

		helpers.RespondWithJSON(w, http.StatusCreated, &comment)

	}

}

func (ch *CommentHandler) DeleteByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		err := ch.CMRepo.DeleteByID(vars["id"])
		if err != nil {
			ch.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Comment hasn't been deleted")
			return
		}

		res := fmt.Sprintf("id: %s", vars["id"])

		helpers.RespondWithJSON(w, http.StatusOK, res)

	}
}

func (ch *CommentHandler) UpdateByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			ch.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var comment domain.Comment
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&comment); err != nil {
			ch.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		comment.ID = uint(id)

		if err := ch.CMRepo.UpdateByID(&comment); err != nil {
			ch.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, comment)
	}
}
