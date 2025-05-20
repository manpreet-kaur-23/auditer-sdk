package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

// SendKafkaMessage sends a message to the specified Kafka topic.
func SendKafkaMessage(topic, msg string) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		return err
	}
	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	_, _, err = producer.SendMessage(message)
	log.Println("Done")
	return err
}
