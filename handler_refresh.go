package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/riyadvr/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find refresh token", err)
		return
	}

	fmt.Printf("refresh token: %v\n", refreshToken)

	refreshTokenRow, err := cfg.db.GetRefreshTokenRow(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find refresh token in the database", err)
		return
	}

	if refreshTokenRow.ExpiresAt.Before(time.Now()) {
		respondWithError(w, http.StatusUnauthorized, "Refresh token expired", err)
		return
	}

	if refreshTokenRow.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "Refresh token revoked", err)
		return
	}

	newAccessToken, err := auth.MakeJWT(refreshTokenRow.UserID, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't refresh the access token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Token: newAccessToken,
	})

}
