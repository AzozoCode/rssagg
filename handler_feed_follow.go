package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/azozocode/rssagg/internal/database"
	"github.com/go-chi/chi"
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

func (apiCfg *apiConfig) handlerGetUserFeedFollowById(w http.ResponseWriter, r *http.Request, user database.User) {

	follow_feed, err := apiCfg.DB.GetFeedFollowByUserID(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving user feed follow: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserFeedFollowToFeedFollows(follow_feed))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feedIdStr := chi.URLParam(r, "feed_id")

	feedId, err := uuid.Parse(feedIdStr)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing feed id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedId,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error deleting feed follow: %v", err))
		return
	}

	respondWithJSON(w, 200, fmt.Sprintf("Feed follow with id %s deleted successfully.", feedIdStr))
}
