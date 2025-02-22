package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/azozocode/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User) {

	type Parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := Parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	follow_feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		FeedID:   params.FeedID,
		ID:       uuid.New(),
		UserID:   user.ID,
		CreateAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user feed follow: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseUserFeedFollowToFeedFollow(follow_feed))
}
