package shared

import "net/http"

func WriteError(w http.ResponseWriter){
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}