package main

import (
	githubclient "dev/github-fav-language/github-client"
	"dev/github-fav-language/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Init router
	router := mux.NewRouter()

	//INSTANTIATE THE HANDLER
	h := handlers.Handler{
		GithubClient: &githubclient.LangugeClient{
			Doer: http.DefaultClient,
		},
	}

	// Route Handlers / Endpoints
	router.HandleFunc("/favLanguage/{userName}", h.GetFavLanguage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
