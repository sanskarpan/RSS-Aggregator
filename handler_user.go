package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/itsemadbattal/rss-aggregator/internal/auth"
	"github.com/itsemadbattal/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	//parsing the request body
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	// when decoding JSON
	//we have to decode JSON into a pointer of the struct because we want the changes to be reflected outside of the function as well
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v .", err))
		return
	}

	//we have to pass the context first then the actual params struct
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnot create user: %v.", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnot get posts: %v, ", err))
	}

	respondWithJSON(w, 200, databasePostsToPosts(posts))
}

func (apiCfg *apiConfig) handlerGetUserByName(w http.ResponseWriter, r *http.Request) {
	name, err := auth.GetName(r.Header)

	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error: %s", err))
		return
	}

	user, err := apiCfg.DB.GetUserByName(r.Context(), name)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnot get user: %s", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
