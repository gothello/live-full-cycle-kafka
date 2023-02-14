package akafka

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consume(topics []string, server string, msgs chan *kafka.Message) error {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          "marcelo_dev",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return err
	}

	consumer.SubscribeTopics(topics, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			return err
		}

		msgs <- msg
	}
}

func Producer(topics []string, server, msg interface{}) error {
	publish, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          "marcelo_dev",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return err
	}

	defer publish.Close()

	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	for _, t := range topics {
		err := publish.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &t, Partition: kafka.PartitionAny},
			Value:          body,
		}, nil)

		if err != nil {
			return err
		}
	}

	return nil
}
