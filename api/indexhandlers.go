package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func createIndex(writer http.ResponseWriter, httpRequest *http.Request) {
	type Request struct {
		Name string `json:"name"`
	}

	var request Request

	errorStatus := tryDecode(&request, httpRequest)

	if errorStatus != nil {
		renderer.JSON(writer, http.StatusOK, errorStatus)
		return
	}

	// todo create the index

	response := Status{
		Status:  "200",
		Message: fmt.Sprintf("Index '%s' created successfully", request.Name),
	}

	renderer.JSON(writer, http.StatusOK, response)
}

func deleteIndex(writer http.ResponseWriter, httpRequest *http.Request) {
	type Request struct {
		Name string `json:"name"`
	}

	var request Request

	errorStatus := tryDecode(&request, httpRequest)

	if errorStatus != nil {
		renderer.JSON(writer, http.StatusOK, errorStatus)
		return
	}

	// todo create the index

	response := Status{
		Status:  "200",
		Message: fmt.Sprintf("Index '%s' deleted successfully", request.Name),
	}

	renderer.JSON(writer, http.StatusOK, response)
}

func tryDecode(value interface{}, request *http.Request) *Status {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(value)

	if err != nil {
		log.Println(err.Error())
		response := Status{
			Status:  "400",
			Message: "Malformed json",
		}

		return &response
	}

	return nil
}
