package api

import (
	"github.com/pressly/chi"
	"net/http"
)

func indexRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", createIndex)
	r.Delete("/", deleteIndex)
	return r
}
