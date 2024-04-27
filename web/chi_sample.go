package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	"github.com/unrolled/render"
)

var rnd *render.Render

type resp struct {
	Message string
}

func RunChiServer(host string, port string) {
	r := chi.NewRouter()

	rnd = render.New()

	// r.Use(middleware.Logger)
	r.Get("/", homeController)
	r.Get("/time", timeController)
	r.Get("/admin", adminController)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r))
}

func homeController(w http.ResponseWriter, r *http.Request) {
	if err := rnd.JSON(w, http.StatusOK, resp{Message: "Hello"}); err != nil {
		log.Println("homeController error:", err)
	}
}

func timeController(w http.ResponseWriter, r *http.Request) {
	if err := rnd.JSON(w, http.StatusOK, resp{Message: time.Now().String()}); err != nil {
		log.Println("timeController error:", err)
	}
}

func adminController(w http.ResponseWriter, r *http.Request) {
	if err := rnd.HTML(w, http.StatusOK, "Goodbye", nil); err != nil {
		log.Println("timeController error:", err)
	}
}
