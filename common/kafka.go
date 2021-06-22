package common

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo/v4"
	"log"
)

var topic = "test-topic"

var consumerConfig = &kafka.ConfigMap{
	"bootstrap.servers": "localhost:9091",
	"group.id":          "test_group",
	"auto.offset.reset": "earliest",
}
var producerConfig = &kafka.ConfigMap{
	"bootstrap.servers": "localhost:9091",
}

func KafkaReceiver() {
	consumer, err := kafka.NewConsumer(consumerConfig)
	if err != nil {
		log.Println("Consumer connection error...!! ------ ", err.Error())
	}
	consumer.Subscribe(topic, nil)

	for {
		message, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Println("Consumer error...!! ------ ", err.Error())
		}
		log.Println(string(message.Key), " : ", string(message.Value), " !!!")
	}
}

func SendMessage(ctx echo.Context) error {
	name := ctx.Param("name")
	producer, err := kafka.NewProducer(producerConfig)
	if err != nil {
		log.Println(err.Error())
	}
	produceError := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(name),
		Value:          []byte("My Name is " + name),
	}, nil)
	if produceError != nil {
		return ErrorResponse(ctx, err.Error(), "send message error...!!")
	}
	producer.Flush(1000)
	return SuccessResponse(ctx, "Done", name)
}
