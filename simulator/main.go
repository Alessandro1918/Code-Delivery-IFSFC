package main

import (
	"fmt"
	// route2 "github.com/Alessandro1918/Code-Delivery-IFSFC/app/route"
	kafka2 "github.com/Alessandro1918/Code-Delivery-IFSFC/app/kafka"
	"github.com/Alessandro1918/Code-Delivery-IFSFC/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

	//test route simulator
	// route := route2.Route{ID: "1", ClientID: "1"}
	// route.LoadPositions()
	// positions, _ := route.ExportJsonPositions()	//list of lat/long positions, error ignored
	// fmt.Println(positions[1])										//one pair of lat/long positions

	//test produce topic
	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("Olá", "readtest", producer)
	// for {
	// 	_ = 1
	// }

	//test consume topic
	// msgChan := make(chan *ckafka.Message)
	// consumer := kafka.NewKafkaConsumer(msgChan)
	// go consumer.Consume()				//start consuming topics in an async way in a new thread
	// for msg := range msgChan {
	// 	fmt.Println(string(msg.Value))
	// }

	//test everything
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()				//start consuming topics in an async way in a new thread
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)		//start producing topics in an async way in a new thread
	}
}