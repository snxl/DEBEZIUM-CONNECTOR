package client

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/snxl/DEBEZIUM-CONNECTOR/utils"
)

func CreateKafkaConsumer() *kafka.Consumer {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9094",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	return consumer
}

func SubscribeTopic(consumer *kafka.Consumer) {
	consumer.Subscribe("Name", nil)

	fmt.Println("Subscribed to topics")
}

func ReadTopicMessages(consumer *kafka.Consumer) string {

	var message string

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {

			fmt.Printf("Topic: %s \n", *msg.TopicPartition.Topic)

			b := []byte(msg.Value)
			b, _ = utils.Prettyprint(b)

			fmt.Printf("%s\n", b)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	wait.Done()
	closeConsumer(consumer)

	return message
}

func closeConsumer(consumer *kafka.Consumer) {
	consumer.Close()
}
