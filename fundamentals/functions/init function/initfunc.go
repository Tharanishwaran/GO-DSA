package main

import (
	"database/sql"
	"log"
)




var database *sql.DB

func init() {
    // Runs automatically before main function
    // Used for initial setup, database connections, etc.
    var err error
    database, err = sql.Open("mysql", "connection_string")
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    // Database is already initialized
}