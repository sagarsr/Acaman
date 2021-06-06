package api

import (
	"Acaman/models"
	"Acaman/utils"
	"encoding/json"
	"net/http"
)

// RegisterUser is used to register a new user.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := user.AddUser(utils.Connection); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))

	}
}
