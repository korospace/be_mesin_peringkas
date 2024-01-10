package routes

import (
	"be_mesin_penerjemah/controllers"

	"github.com/gorilla/mux"
)

func PenerjemahRoutes(r *mux.Router) {
	router := r.PathPrefix("/penerjemah").Subrouter()

	router.HandleFunc("/run", controllers.RunPenerjemah).Methods("POST", "OPTIONS")
}
