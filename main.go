package main

import (
	"github.com/snxl/DEBEZIUM-CONNECTOR/cmd/client"
)

func main() {

	consumer := client.CreateKafkaConsumer()
	client.ConnectorHealthCheck()

	client.SubscribeTopic(consumer)
	client.ReadTopicMessages(consumer)

}
