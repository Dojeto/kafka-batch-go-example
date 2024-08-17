package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/dojeto/kafka-batch-go-example/model"
	"github.com/gin-gonic/gin"
)

var (
	topic     = "userData"
	brokerUrl = []string{"locahost:9092"}
	config    = sarama.NewConfig()
)

func main() {
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	r := gin.Default()

	k, err := NewKafkaClient(brokerUrl, config)

	r.POST("/", func(ctx *gin.Context) {

		if err != nil {
			log.Fatalf("Error while Creating Kafka Client %s", err)
			return
		}

		data, err := ReadCsv()

		if err != nil {
			log.Fatalf("Error while Reading Csv Files %s", err)
		}

		for _, value := range data {
			x := model.User{
				Name:     value.Name,
				Email:    value.Email,
				Password: value.Password,
			}
			y, err := json.Marshal(x)
			if err != nil {
				log.Fatal(err)
			}
			k.PushToKafka(topic, y)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"sucess": true,
			"data":   "Data Has Been Added SuccessFully",
		})
	})
	r.Run()
}
