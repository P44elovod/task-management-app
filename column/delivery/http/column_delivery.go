package columnhttpdelivery

import (
	"encoding/json"
	"net/http"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ColumnHandler struct {
	CUsecase domain.ColumnUsecase
	logger   *logrus.Logger
}

func New(r *mux.Router, log *logrus.Logger, cu domain.ColumnUsecase) {
	handler := &ColumnHandler{
		CUsecase: cu,
	}

	r.HandleFunc("/column/new", handler.Create()).Methods("POST")

}

func (c *ColumnHandler) Create() http.HandlerFunc {
	var column domain.Column

	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&column); err != nil {
			c.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		err := c.CUsecase.CreateColumn(&column)
		if err != nil {
			c.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Column not created")
			return
		}
		defer r.Body.Close()

		helpers.RespondWithJSON(w, http.StatusCreated, &column)

	}

}
