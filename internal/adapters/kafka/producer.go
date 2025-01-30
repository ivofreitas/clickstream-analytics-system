package kafka

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ivofreitas/clickstream-analytics-system/internal/domain"
)

type KafkaProducer struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaProducer(broker, topic string) *KafkaProducer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	return &KafkaProducer{producer: p, topic: topic}
}

func (kp *KafkaProducer) PublishEvent(event domain.Event) error {
	data, _ := json.Marshal(event)
	return kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kp.topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)
}
