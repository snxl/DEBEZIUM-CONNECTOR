package main

import (
	"database/sql"
	"fmt"
	"github.com/snxl/DEBEZIUM-CONNECTOR/cmd/client"
	"github.com/snxl/DEBEZIUM-CONNECTOR/config"
	"sync"
)

var wait sync.WaitGroup
var db *sql.DB

func main() {
	db, err := sql.Open(config.PostgresDriver, config.DataSourceName)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected")
	}
	defer db.Close()

	client.CreateTable(db)

	consumer := client.CreateKafkaConsumer()
	client.ConnectorHealthCheck()

	client.SubscribeTopic(consumer)

	wait.Add(2)
		go client.InsertName(db)
		go client.ReadTopicMessages(consumer)
	wait.Wait()

}
