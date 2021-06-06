package api

import (
	"net/http"
)

//HomeLink Request returning home page
func HomeLink(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome home!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Navigate to /registeruser"}`))
}
