package main

import (
	"dev/github-fav-language/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Init router
	router := mux.NewRouter()

	// Route Handlers / Endpoints
	router.HandleFunc("/favLanguage", handlers.GetFavLanguage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
