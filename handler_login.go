package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/riyadvr/chirpy/internal/auth"
	"github.com/riyadvr/chirpy/internal/database"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email            string `json:"email"`
		Password         string `json:"password"`
		ExpiresInSeconds int    `json:"expires_in_seconds"`
	}

	type response struct {
		User
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couln't decode parameter", err)
		return
	}

	dbUser, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}
	match, err := auth.CheckPasswordHash(params.Password, dbUser.HashedPassword)
	if err != nil || !match {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	// expirationTime := time.Hour
	// if params.ExpiresInSeconds > 0 && params.ExpiresInSeconds < 3600 {
	// 	expirationTime = time.Duration(params.ExpiresInSeconds) * time.Second
	// }

	accessToken, err := auth.MakeJWT(dbUser.ID, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "JWT creation failed", err)
		return
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Refresh token creation failed", err)
		return
	}

	_, err = cfg.db.AddRefreshToken(r.Context(), database.AddRefreshTokenParams{
		Token:     refreshToken,
		UserID:    dbUser.ID,
		ExpiresAt: time.Now().Add(60 * 24 * time.Hour),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't add refresh token to the database", err)
		return
	}

	apiUser := response{}
	if match {
		apiUser.ID = dbUser.ID
		apiUser.CreatedAt = dbUser.CreatedAt
		apiUser.UpdatedAt = dbUser.UpdatedAt
		apiUser.Email = dbUser.Email
		apiUser.Token = accessToken
		apiUser.RefreshToken = refreshToken
		apiUser.IsChirpyRed = dbUser.IsChirpyRed
	}
	respondWithJSON(w, http.StatusOK, apiUser)
}
