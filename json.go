package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	//errors 500 and above are from the server end thats why we only check that
	//we dont care about slient side errors (< 500)
	if code > 499 {
		log.Printf("Responding with %v error: %v", code, msg)
	}

	type errResponse struct {
		Error string `json:"error"` //=> {"error" : "msg"}
	}

	respondWithJSON(w, code, errResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//converting the payload to json
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	//after the marshaling succeeded, we tell that we are writing with JSON format
	w.Header().Add("Content-Type", "application/json")
	//we set a status code of the passed in response code to indicate success
	w.WriteHeader(code)
	//we write the data
	w.Write(data)
}
