package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/azozocode/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUserFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type Parameters struct {
		Name string `json:"name"`
		URl  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := Parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:       uuid.New(),
		Name:     params.Name,
		CreateAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
		Url:      params.URl,
		UserID:   user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseUserFeedToFeed(feed))
	return

}

func (apiCfg *apiConfig) handlerGetUserFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting user feeds: %v", err))
		return
	}
	respondWithJSON(w, http.StatusOK, feeds)
	return
}
