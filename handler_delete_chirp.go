package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/riyadvr/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
	chirpID := r.PathValue("chirpID")
	chirpUUID, err := uuid.Parse(chirpID)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	accessToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}

	userId, err := auth.ValidateJWT(accessToken, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	chirp, err := cfg.db.GetChirpById(r.Context(), chirpUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}

	if chirp.UserID != userId {
		respondWithError(w, http.StatusForbidden, "forbidden: cannot delete another user's chirp", nil)
		return
	}

	if err := cfg.db.DeleteChirpById(r.Context(), chirpUUID); err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
