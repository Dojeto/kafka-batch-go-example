package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

type KafkaClient struct {
	conn sarama.SyncProducer
}

func NewKafkaClient(broker []string, config *sarama.Config) (*KafkaClient, error) {
	conn, err := sarama.NewSyncProducer(broker, config)
	if err != nil {
		return nil, err
	}
	return &KafkaClient{
		conn: conn,
	}, nil
}

func (k *KafkaClient) PushToKafka(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := k.conn.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Message Stored in Topic (%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}
