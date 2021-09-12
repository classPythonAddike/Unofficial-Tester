package main

import (
	"log"
	"net/http"
	"os"
	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	// "github.com/go-chi/httprate"
)

func main() {
	/*
		r := chi.NewRouter()

		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(httprate.LimitByIP(10, 1*time.Minute))

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			WriteMessage(&w, message_from_yan, http.StatusOK)
		})

		r.Get("/testing", RunFile)
	*/

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			WriteMessage(&w, message_from_yan)
		},
	)

	http.Handle(
		"/run-test",
		Logger(http.HandlerFunc(RunFile)),
	)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
