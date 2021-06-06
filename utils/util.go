package utils

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v4"
)

//Connection global
var Connection *pgx.Conn

//SetConnection sets connection to database cross package
func SetConnection(Con *pgx.Conn) {
	Connection = Con

}

//RespondWithError uses
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

//RespondWithJSON uses
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
