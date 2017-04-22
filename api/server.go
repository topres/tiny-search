package api

import (
	"log"
	"net/http"

	"github.com/pressly/chi"
	//"github.com/pressly/chi/middleware"
	"github.com/unrolled/render"
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
	router.Mount("/index", homeRouter())

	log.Println("Starting HTTP server on port " + settings.Port)

	http.ListenAndServe(":"+settings.Port, router)
}
