package main

import "net/http"

//BuildJSONResponse decorate the http responseWriter
func BuildJSONResponse(w http.ResponseWriter, code int) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code) // unprocessable entity
	return w
}
