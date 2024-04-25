package web

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RunGoChiServer() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", homeChi)

	log.Fatal(http.ListenAndServe(":8000", r))
}

func homeChi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().String() + "\n"))
}
