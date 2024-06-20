package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code >= 500 {
		log.Printf("Responding with error 5XX error: %s", msg)
	}

	type errorResponde struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorResponde{
		Error: msg,
	})

}
