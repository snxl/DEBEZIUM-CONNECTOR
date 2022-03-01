package main

import (
	"fmt"
	"github.com/snxl/DEBEZIUM-CONNECTOR/cmd/client"
)

func main() {

	consumer := client.CreateKafkaConsumer()
	client.ConnectorHealthCheck()

	client.SubscribeTopic(consumer)
	client.ReadTopicMessages(consumer)

	fmt.Scan()
}
