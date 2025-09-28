package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		log.Printf("current server hits %v", cfg.fileserverHits.Load())
		next.ServeHTTP(w, r)
	})
}
