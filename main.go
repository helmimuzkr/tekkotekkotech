package main

import (
	"log/slog"
	"net/http"
)

const (
	port = ":8181"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("ini log", "method", r.Method, "url", r.URL)
		http.Redirect(w, r, "https://youtu.be/OszmhdV0_cg?si=9WaIQ5hNEWTWiltR", http.StatusSeeOther)
	})

	slog.Info("starting server on", "port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}
}
