package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(httprate.LimitByIP(10, 1*time.Minute))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		WriteMessage(&w, message_from_yan)
		w.WriteHeader(http.StatusOK)
	})

	r.Get("/run-test", RunFile)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Printf("Listening on port %v\n", port)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
