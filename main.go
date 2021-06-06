package main

import (
	view "Acaman/api" //Route Handling
	"Acaman/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

//Connection global
var Connection *pgx.Conn

func main() {
	fmt.Println(os.Getenv("DATABASE_URL"))
	Connection, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Print("error!")
	}
	utils.SetConnection(Connection)
	defer Connection.Close(context.Background())
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", view.HomeLink)
	router.HandleFunc("/register", view.RegisterUser)
	router.HandleFunc("/attendance", view.Attendance)
	router.HandleFunc("/assignments", view.Assignments)
	log.Fatal(http.ListenAndServe(":8080", router))
}
