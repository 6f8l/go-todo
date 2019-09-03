package handler

import (
	"net/http"

	"../db"
	"../service"
)

type todoHandler struct {
	sample *db.Sample
}

func (handler *todoHandler) GetSamples(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handler.samples)

	todoList, err := service.GeaAll(ctx)

	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseOk(w, todoList)
}
