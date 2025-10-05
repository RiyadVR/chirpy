package main

import (
	"encoding/json"
	"net/http"

	"github.com/riyadvr/chirpy/internal/auth"
	"github.com/riyadvr/chirpy/internal/database"
)

func (cfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	accessToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find jwt", err)
		return
	}

	userId, err := auth.ValidateJWT(accessToken, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params parameters
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode JSON", err)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
	}

	dbParams := database.UpdateEmailAndPasswordParams{
		ID:             userId,
		Email:          params.Email,
		HashedPassword: hashedPassword,
	}

	dbUpdatedUser, err := cfg.db.UpdateEmailAndPassword(r.Context(), dbParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "User creation failed", err)
		return
	}

	apiUser := User{
		ID:        dbUpdatedUser.ID,
		CreatedAt: dbUpdatedUser.CreatedAt,
		UpdatedAt: dbUpdatedUser.UpdatedAt,
		Email:     dbParams.Email,
	}

	respondWithJSON(w, http.StatusOK, apiUser)
}
