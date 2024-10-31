package main

import (
	"log"
	"net/http"

	"github.com/MdHasib01/gorillMux_postgres_api/config"
	"github.com/MdHasib01/gorillMux_postgres_api/db"

	"github.com/MdHasib01/gorillMux_postgres_api/routes"

	"github.com/gorilla/mux"
)

func main() {
    config.LoadEnv()  
    db.Init()         

    r := mux.NewRouter()
    routes.RegisterRoutes(r)

    log.Println("Server is starting at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
