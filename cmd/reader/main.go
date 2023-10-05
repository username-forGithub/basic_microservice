package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	fmt.Println("hhhhhhhhhh")

	config := &kafka.ConfigMap{
		"bootstrap.servers": "uzumdeliv-kafka-1:29092",
		"group.id":          "my-consumer-group",
		"auto.offset.reset": "earliest",
	}

	fmt.Println(config)
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Println("failed to creat consumer")
		return
	}
	fmt.Println("zzzzz")

	topics := []string{"mytopic"}
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Println("failed to send topics")
		return
	}

	log.Println("жду сообщений--- ")
	//fmt.Println("xxx")
	//for {
	//	msg, err := consumer.ReadMessage(-1)
	//	if err == nil {
	//		log.Println(string(msg.Value))
	//		return
	//	} else {
	//		fmt.Println("error readmsg")
	//	}
	//}
}
