package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	//we respond with an empty struct cause we only care about the status code to indicate success or failure
	respondWithJSON(w, 200, struct{}{})
}
