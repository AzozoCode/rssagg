package main

import (
	"fmt"
	"net/http"

	"github.com/azozocode/rssagg/internal/auth"
	"github.com/azozocode/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) authMiddleware(handler authedHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		api_key, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), api_key)

		if err != nil {
			respondWithError(w, http.StatusNotFound, fmt.Sprintf("User not found: %v", err))
			return
		}

		handler(w, r, user)
	}
}
