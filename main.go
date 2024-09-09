package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"simpleBank/tutorial"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://nei:54321@localhost:5432/control_system_db?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	ctx := context.Background()
	queries := tutorial.New(conn)

	list, err := queries.ListAccounts(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("list", list)

	arg := tutorial.CreateAccountParams{
		Owner:    "Tom & Jerry",
		Balance:  12000000,
		Currency: "USD",
	}

	account, err := queries.CreateAccount(ctx, arg)

	fmt.Println("account ", account)

	price, err := queries.TotalBalance(ctx)
	fmt.Println("total balace", price)

	defer conn.Close()
}
