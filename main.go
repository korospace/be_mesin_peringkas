package main

import (
	"be_mesin_penerjemah/middleware"
	"be_mesin_penerjemah/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	// Global Middleware
	router.Use(middleware.CorsMiddleware)

	// Routes Registration
	routes.PenerjemahRoutes(router)

	log.Println("server running at http://localhost:8282")
	http.ListenAndServe(":8282", router)
}
