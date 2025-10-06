package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
)

// this handler needs to refactor the duplicate code later
func (cfg *apiConfig) handlerListChirps(w http.ResponseWriter, r *http.Request) {

	authorID := r.URL.Query().Get("author_id")
	sortvalue := r.URL.Query().Get("sort") // sorting by value desc and asc, asc is be default by sql query

	if authorID != "" { // this block means the author_id is present the query parameters
		authorUUID, err := uuid.Parse(authorID)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid UUID", err)
			return
		}

		dbChirps, err := cfg.db.GetChirpByUserID(r.Context(), authorUUID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
			return
		}

		chirps := []Chirp{}
		for _, dbChirp := range dbChirps {
			chirps = append(chirps, Chirp{
				ID:        dbChirp.ID,
				CreatedAt: dbChirp.CreatedAt,
				UpdatedAt: dbChirp.UpdatedAt,
				UserID:    dbChirp.UserID,
				Body:      dbChirp.Body,
			})
		}

		if sortvalue == "desc" {
			sort.Slice(chirps, func(i, j int) bool { return chirps[i].CreatedAt.After(chirps[j].CreatedAt) })
		}

		respondWithJSON(w, http.StatusOK, chirps)
		return

	}

	dbChirps, err := cfg.db.ListChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			UserID:    dbChirp.UserID,
			Body:      dbChirp.Body,
		})
	}

	if sortvalue == "desc" {
		sort.Slice(chirps, func(i, j int) bool { return chirps[i].CreatedAt.After(chirps[j].CreatedAt) })
	}

	respondWithJSON(w, http.StatusOK, chirps)
}
