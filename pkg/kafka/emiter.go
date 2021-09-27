package kafka

import (
	"encoding/json"

	"bitbucket.org/josecaceresatencora/logistics/pkg/bus"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	stdkafka "github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

type (
	Emiter struct {
		Producer *stdkafka.Producer
		Topic    string
	}
)

func (e *Emiter) Emit(event string, payload []byte) error {
	evt, err := json.Marshal(bus.Event{
		Event:   event,
		Payload: payload,
	})
	if err != nil {
		log.Error(err)
	}

	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &e.Topic,
			Partition: kafka.PartitionAny,
		},
		Value: evt,
	}

	return e.Producer.Produce(&msg, nil)
}
