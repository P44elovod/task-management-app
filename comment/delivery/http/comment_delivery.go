package commenthttpdelivery

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
	"github.com/gorilla/mux"
)

type CommentHandler struct {
	CMUsecase domain.CommentUseCase
}

func New(r *mux.Router, cmu domain.CommentUseCase) {
	handler := &CommentHandler{
		CMUsecase: cmu,
	}

	r.HandleFunc("/comment/new", handler.Create()).Methods("POST")
}

func (ch *CommentHandler) Create() http.HandlerFunc {
	var comment domain.Comment

	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&comment); err != nil {
			log.Print(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		}

		err := ch.CMUsecase.CreateComment(&comment)
		if err != nil {
			log.Print(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Comment not created")
		}
		defer r.Body.Close()

		helpers.RespondWithJSON(w, http.StatusCreated, &comment)

	}

}
