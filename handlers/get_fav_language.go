package handlers

import (
	"fmt"
	"net/http"
)

func GetFavLanguage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("fav language")

	// param := mux.Vars(r)["id"]
	// id, err := strconv.ParseUint(param, 10, 64)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	// post, err := model.GetPost(id)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	// json.NewEncoder(w).Encode(post)

}
