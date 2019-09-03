package handler

import (
	"encoding/json"
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

func responseOk(w http.ResponseWrite, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(body)
}

func responseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}
