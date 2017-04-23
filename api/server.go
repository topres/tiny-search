package api

import (
	"github.com/pressly/chi"
	"github.com/unrolled/render"
	"log"
	"net/http"
)

var renderer *render.Render

func StartServer(settings AppSettings) {

	renderer = render.New(render.Options{
		IndentJSON: true,
		Layout:     "layout",
	})

	router := chi.NewRouter()

	router.Use(AppMiddleware(settings))

	// routes
	router.Mount("/index", indexRouter())

	log.Println("Starting HTTP server on port " + settings.Port)

	http.ListenAndServe(":"+settings.Port, router)
}
