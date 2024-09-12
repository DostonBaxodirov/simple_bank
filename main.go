package main

import (
	"database/sql"
	"log"
	"simpleBank/api"
	"simpleBank/tutorial"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://nei:54321@localhost:5432/control_system_db?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := tutorial.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
