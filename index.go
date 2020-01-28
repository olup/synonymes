package main

import (
	"encoding/json"
	"net/http"

	handler "github.com/olup/api/handler"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	dictionary := handler.Open()
	word := r.URL.Query()["word"]
	entry := dictionary[word[0]]
	JSON, _ := json.Marshal(entry)
	w.Header().Set("Cache-Control", "s-maxage=3155690, stale-while-revalidate")
	w.Header().Set("Content-Type", "application/json")
	w.Write(JSON)
}
