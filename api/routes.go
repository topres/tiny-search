package api

import (
	"github.com/pressly/chi"
	"net/http"
)

func homeRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeIndex)
	return r
}
