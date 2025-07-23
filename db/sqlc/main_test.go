package sqlc

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
	"udemy-course/utils"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatalln("Cannot load configurations!")
	}
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("We cannot connect to database", err)
	}
	testQueries = New(testDb)

	os.Exit(m.Run())
}
