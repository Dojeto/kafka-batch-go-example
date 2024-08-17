package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/dojeto/kafka-batch-go-example/consumer/utils"
	"github.com/dojeto/kafka-batch-go-example/model"
)

var (
	timer *time.Timer

	config    = sarama.NewConfig()
	topic     = "userData"
	brokerUrl = []string{"localhost:9092"}
	duration  = 2 * time.Second
	data      []model.User

	res  *utils.Response
	stop = make(chan int)
)

const (
	batchSize int = 5
)

func init() {
	ConnectToDb()
	DB.AutoMigrate()
}

func main() {
	config.Consumer.Return.Errors = true
	config.Consumer.MaxWaitTime = 3 * time.Second

	fmt.Println(brokerUrl)

	conn, err := sarama.NewConsumer(brokerUrl, config)
	if err != nil {
		log.Fatalln("Connection Failed :", err)
		return
	}
	fmt.Println("Consumer Is Ready")
	defer conn.Close()

	consumer, err := conn.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Printf("Failed to consume partition: %v", err)
		return
	}
	defer consumer.Close()

	ch := consumer.Messages()

	go func() {
		for x := range ch {
			if len(data) >= batchSize {
				DB.Create(&data)
				data = []model.User{}
			}
			err = json.Unmarshal(x.Value, res)
			data = append(data, model.User{
				Name:     string(res.Name),
				Email:    string(res.Email),
				Password: string(res.Password),
			})
			DeBouncing(duration, &data, stop)
		}
	}()
	<-stop
	close(stop)
}

// DeBouncing is a Famouse Technique

func DeBouncing(d time.Duration, data *[]model.User, stop chan int) {
	if timer != nil {
		timer.Stop()
	}
	timer = time.AfterFunc(d, func() {
		DB.Create(data)
		log.Println("Task Completed !!")
		stop <- 2
	})
}
