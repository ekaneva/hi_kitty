package main

import (
	"net/http"
)

type requestHandler struct {
}

func (handler *requestHandler) helloMessage(writer http.ResponseWriter, request *http.Request) {
    defer request.Body.Close()
    writer.WriteHeader(http.StatusOK)
    writer.Write([]byte(`Hello Cookiecutter`))
}

func (handler *requestHandler) sampleMessage(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	// todo: implement handler logic

	err := error(nil)
	switch err {
	case nil:
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(`{"output":"Sample output"}`))
	default:
		writer.WriteHeader(http.StatusInternalServerError)
	}

}
