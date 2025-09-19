package main

import (
	"fmt"
	"net/http"

	"github.com/itsemadbattal/rss-aggregator/internal/auth"
	"github.com/itsemadbattal/rss-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			//403 is permission error
			respondWithError(w, 403, fmt.Sprintf("Auth error: %s", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldnot get user by api_key: %s", err))
			return
		}

		handler(w, r, user)
	}
}
