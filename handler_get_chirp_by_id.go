package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirpById(w http.ResponseWriter, r *http.Request) {
	chirpID := r.PathValue("chirpID")
	chirpUUID, err := uuid.Parse(chirpID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	dbChirp, err := cfg.db.GetChirpById(r.Context(), chirpUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't retrieve the chirp", err)
		return
	}

	apiChirp := Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserID:    dbChirp.UserID,
	}

	respondWithJSON(w, http.StatusOK, apiChirp)
}
