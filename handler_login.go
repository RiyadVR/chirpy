package main

import (
	"encoding/json"
	"net/http"

	"github.com/riyadvr/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couln't decode parameter", err)
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
	apiUser := User{}
	if match {
		apiUser.ID = dbUser.ID
		apiUser.CreatedAt = dbUser.CreatedAt
		apiUser.UpdatedAt = dbUser.UpdatedAt
		apiUser.Email = dbUser.Email
	}
	respondWithJSON(w, http.StatusOK, apiUser)
}
