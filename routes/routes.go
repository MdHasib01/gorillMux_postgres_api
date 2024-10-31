package routes

import (
	"github.com/MdHasib01/gorillMux_postgres_api/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
    r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
}
