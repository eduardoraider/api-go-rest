package routes

import (
	"api-go-rest/controllers"
	"api-go-rest/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func HandleRequest() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
	r.Use(middleware.ContentTypeMiddleware)
	r.Get("/", controllers.Home)
	r.Get("/api/personalities", controllers.GetPersonalities)
	r.Get("/api/personalities/{id:[0-9]+}", controllers.GetPersonality)
	r.Post("/api/personalities", controllers.CreatePersonality)
	r.Delete("/api/personalities/{id:[0-9]+}", controllers.DeletePersonality)
	r.Put("/api/personalities/{id:[0-9]+}", controllers.EditPersonality)
	log.Fatal(http.ListenAndServe(":8000", r))
}
