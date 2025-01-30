package kafka

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ivofreitas/clickstream-analytics-system/internal/app"
	"github.com/ivofreitas/clickstream-analytics-system/internal/domain"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
	service  *app.EventService
}

func NewKafkaConsumer(broker, topic, group string, service *app.EventService) *KafkaConsumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          group,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	c.SubscribeTopics([]string{topic}, nil)
	return &KafkaConsumer{consumer: c, service: service}
}

func (kc *KafkaConsumer) Start() {
	for {
		msg, err := kc.consumer.ReadMessage(-1)
		if err == nil {
			var event domain.Event
			json.Unmarshal(msg.Value, &event)
			kc.service.TrackEvent(event)
		}
	}
}
