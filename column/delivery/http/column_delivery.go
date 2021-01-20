package columnhttpdelivery

import (
	"encoding/json"
	"net/http"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
	"github.com/gorilla/mux"
)

type ColumnHandler struct {
	CUsecase domain.ColumnUsecase
}

func New(r *mux.Router, pu domain.ColumnUsecase) {
	handler := &ColumnHandler{
		CUsecase: pu,
	}

	r.HandleFunc("/column/new", handler.Create()).Methods("POST")

}

func (c *ColumnHandler) Create() http.HandlerFunc {
	var column domain.Column

	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&column); err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		}

		err := c.CUsecase.CreateColumn(&column)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Column not created")
		}
		defer r.Body.Close()

		helpers.RespondWithJSON(w, http.StatusCreated, &column)

	}

}
