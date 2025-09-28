package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type retrunVals struct {
		Cleaned_body string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
	}

	const maxChirpLength = 140

	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	respondWithJSON(w, http.StatusOK, retrunVals{
		Cleaned_body: cleanChirp(params.Body),
	})
}

func cleanChirp(chirp string) string {
	splittedWords := strings.Split(chirp, " ")

	for i, word := range splittedWords {
		if strings.ToLower(word) == "kerfuffle" {
			splittedWords[i] = "****"
		}
		if strings.ToLower(word) == "sharbert" {
			splittedWords[i] = "****"
		}
		if strings.ToLower(word) == "fornax" {
			splittedWords[i] = "****"
		}
	}
	return strings.Join(splittedWords, " ")

}
