package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"simpleBank/api"
	"simpleBank/tutorial"
	"simpleBank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot connect to config file", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := tutorial.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("Cannot create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
