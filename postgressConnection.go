package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

	//postgresql://postgres:postgres@localhost:5432/GoappDB
	// Connect to the PostgreSQL database
	connStr := "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	//	db, err := sql.Open("pgx", connStr)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to Postgress")
	}

	defer db.Close()

	// Perform a sample query
	rows, err := db.Query("SELECT * FROM \"Trains\"")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	} else {
		fmt.Println("Connection Established")
	}
	defer rows.Close()

	// Iterate through the results
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}
