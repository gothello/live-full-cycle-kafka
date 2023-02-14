package akafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func Consume(topics []string, server string, msgs chan *kafka.Message) error {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          "imersao12-go-esquenta",
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
