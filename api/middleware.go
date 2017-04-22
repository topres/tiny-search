package api

import (
	"context"
	"log"
	"net/http"
)

func AppMiddleware(app AppSettings) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(writer http.ResponseWriter, request *http.Request) {
			ctx := context.WithValue(request.Context(), "app", app)
			writer.Header().Set("x-tiny-searcher-version", "0.0.1")
			next.ServeHTTP(writer, request.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func RecoverMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("%s %s panic: %+v", r.Method, r.RequestURI, err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
