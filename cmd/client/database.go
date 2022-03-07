package client

import (
	"database/sql"
	"github.com/snxl/DEBEZIUM-CONNECTOR/utils"
	"sync"
	"time"
)

var wait sync.WaitGroup

func CreateTable(db *sql.DB) {

	query := `
		CREATE TABLE IF NOT EXISTS "Name" (
			id SERIAL PRIMARY KEY NOT NULL,
			name VARCHAR NOT NULL 
		)
	`

	_, err := db.Query(query)

	if err != nil {
		panic(err)
	}

}

func InsertName(db *sql.DB) {

	for {
		randName := utils.RandSeq(8)

		_, err := db.Exec("INSERT INTO \"Name\" (id, name) VALUES (DEFAULT, $1 )", randName)

		if err != nil {
			panic(err)
		}

		time.Sleep(time.Second * 3)
	}
	wait.Done()
}
