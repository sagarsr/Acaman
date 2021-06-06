package main

import (
	"fmt"

	"github.com/jackc/pgx/v4"
)

//ConnectDatabase uses
func ConnectDatabase(Connection *pgx.Conn) {

	fmt.Println(Connection)

	// sum := 0
	// rows, err := Connection.Query(context.Background(), "select * from faculty;")
	// //fmt.Println(tag, err)
	// for rows.Next() {
	// 	var n int
	// 	var dn string
	// 	var ln string
	// 	err = rows.Scan(&n, &dn, &ln)
	// 	if err != nil {
	// 		fmt.Println("error occurred")
	// 	}
	// 	fmt.Println("ID no:", n, dn, ln)
	// 	sum += n
	// }
	// fmt.Println(sum)
	// row, err := Connection.Query(context.Background(), "Insert into faculty(facid, dname, lname) VALUES ($1,$2,$3)", 8979, "Kayanth", "Parekh")
	// fmt.Println(row, err)
}
