package client

import (
	"bytes"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"io/ioutil"
	"net/http"
)

func CreateKafkaConsumer() *kafka.Consumer {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	return consumer
}

func ConnectorHealthCheck() {
	response, err := http.Get("http://localhost:8083/connectors/tables_connector")
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}
	if response.StatusCode != 200 {
		RegisterConnector()
	}
}

func RegisterConnector() *http.Response {
	plan, _ := ioutil.ReadFile("../../connector/debezium-connector.json")
	response, err := http.Post("http://localhost:8083/connectors/", "application/json", bytes.NewBuffer(plan))
	fmt.Println(bytes.NewBuffer(plan))

	if err != nil {
		panic(err)
	}

	return response
}

func SubscribeTopic(consumer *kafka.Consumer) {
	consumer.Subscribe("postgres.public.Product", nil)

	fmt.Println("Subscribed to product topic")
}

func ReadTopicMessages(consumer *kafka.Consumer) string {

	var message string
	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	closeConsumer(consumer)

	return message
}

func closeConsumer(consumer *kafka.Consumer) {
	consumer.Close()
}
