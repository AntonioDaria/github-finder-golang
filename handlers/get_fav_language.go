package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) GetFavLanguage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// query := r.URL.Query()
	// user_name := query.Get("user_name")
	params := mux.Vars(r)

	language, err := h.GithubClient.GetFavLanguage(r.Context(), params["userName"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(language)
}
