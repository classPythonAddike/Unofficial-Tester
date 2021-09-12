package main

import (
	"log"
	"net/http"

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
			WriteMessage(&w, message_from_yan, http.StatusOK)
		},
	)

	http.Handle(
		"/run-test",
		Logger(http.HandlerFunc(RunFile)),
	)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}