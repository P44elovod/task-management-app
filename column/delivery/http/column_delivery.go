package columnhttpdelivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/P44elovod/task-management-app/domain"
	"github.com/P44elovod/task-management-app/helpers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ColumnHandler struct {
	CUsecase domain.ColumnUsecase
	CRepo    domain.ColumnRepository
	logger   *logrus.Logger
}

func New(r *mux.Router, log *logrus.Logger, cu domain.ColumnUsecase) {
	handler := &ColumnHandler{
		CUsecase: cu,
	}

	r.HandleFunc("/column/new", handler.Create()).Methods("POST")
	r.HandleFunc("/column/{id:[0-9]+}", handler.UpdateByID()).Methods("PUT")
	r.HandleFunc("/column/{id:[0-9]+}/move", handler.UpdatePosiotionByID()).Methods("PUT")
	r.HandleFunc("/column/{id:[0-9]+}", handler.DeleteByID()).Methods("DELETE")

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

func (c *ColumnHandler) UpdateByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			c.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid column ID")
			return
		}

		var column domain.Column
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&column); err != nil {
			c.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		column.ID = uint(id)

		if err := c.CUsecase.Update(&column); err != nil {
			c.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Column not updated")
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, column)
	}
}

func (c *ColumnHandler) UpdatePosiotionByID() http.HandlerFunc {

	type positions struct {
		id       uint
		position uint
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			c.logger.Error(err, id)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid column ID")
			return
		}

		var positionsList []positions
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&positionsList); err != nil {
			c.logger.Error(err)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()

		positionsMap := make(map[uint]uint)
		for i := 0; i < len(positionsList); i++ {
			positionsMap[positionsList[i].id] = positionsList[i].position
		}

		if err := c.CUsecase.UpdatePosition(positionsMap); err != nil {
			c.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Columns postions not updated")
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, positionsList)
	}
}

func (c *ColumnHandler) DeleteByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			c.logger.Error(err, id)
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid column ID")
			return
		}

		if err := c.CUsecase.DeleteByID(uint(id)); err != nil {
			c.logger.Error(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Column hasn't been deleted")
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, id)

	}
}
