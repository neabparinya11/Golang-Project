package queue

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"log"

	"github.com/IBM/sarama"
	"github.com/go-playground/validator/v10"
)

func ConnectProducer(brokerUrls []string, apiKey, secret string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	if apiKey != "" && secret != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = apiKey
		config.Net.SASL.Password = secret
		config.Net.SASL.Mechanism = "PLAIN"
		config.Net.SASL.Handshake = true
		config.Net.SASL.Version = sarama.SASLHandshakeV1
		config.Net.TLS.Enable = true
		config.Net.TLS.Config = &tls.Config{
			InsecureSkipVerify: true,
			ClientAuth: tls.NoClientCert,
		}
	}

	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 3

	producer, err := sarama.NewSyncProducer(brokerUrls, config)
	if err != nil {
		return nil, errors.New("error: Connection failed to kafka")
	}

	return producer, nil
}

func PushMessageWithKeyToQueue(brokerUrls []string, apiKey, secrete, topic, key string, message []byte) error {
	producer, err := ConnectProducer(brokerUrls, apiKey, secrete)
	if err != nil {
		return errors.New("error: Failed to connect to producer")
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
		Key: sarama.StringEncoder(key),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return errors.New("error: Failed to send message")
	}
	log.Printf("Message is store in topic(%s)/ partition(%d) / offset(%d) \n", topic, partition, offset)

	return nil
}

func ConnectConsumer(brokerUrls []string, apiKey, secret string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	if apiKey != "" && secret != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = apiKey
		config.Net.SASL.Password = secret
		config.Net.SASL.Mechanism = "PLAIN"
		config.Net.SASL.Handshake = true
		config.Net.SASL.Version = sarama.SASLHandshakeV1
		config.Net.TLS.Enable = true
		config.Net.TLS.Config = &tls.Config{
			InsecureSkipVerify: true,
			ClientAuth: tls.NoClientCert,
		}
	}

	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 3

	consumer, err := sarama.NewConsumer(brokerUrls, config)
	if err != nil {
		return nil, errors.New("error: Failed to connect to consumer")
	}

	return consumer, nil
}

func DecodeMessage(object any, value []byte) error {
	if err := json.Unmarshal(value, &object); err != nil {
		return errors.New("error: Failed to decode message")
	}

	validate := validator.New()
	if err := validate.Struct(object); err != nil {
		return errors.New("error: Failed to validate message")
	}

	return nil
}