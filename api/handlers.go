package api

import (
	"net/http"
)

func homeIndex(writer http.ResponseWriter, request *http.Request) {
	context := request.Context()
	app, ok := context.Value("app").(AppSettings)

	if !ok {
		http.Error(writer, http.StatusText(422), 422)
		return
	}

	renderer.JSON(writer, http.StatusOK, app)
}
