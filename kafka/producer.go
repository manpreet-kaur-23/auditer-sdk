package kafka

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/manpreet-kaur-23/auditer-sdk/auditer"
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

func SendAuditLogToKafka(topic string, logData auditer.AuditLog) error {
	jsonBytes, err := json.Marshal(logData)
	if err != nil {
		log.Printf("Failed to marshal AuditLog: %v", err)
		return err
	}

	return SendKafkaMessage(topic, string(jsonBytes))
}
