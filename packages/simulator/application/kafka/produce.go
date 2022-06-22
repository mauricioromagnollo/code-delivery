package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	route2 "github.com/x0n4d0/code-delivery-simulator/application/route"
	"github.com/x0n4d0/code-delivery-simulator/infra/kafka"
)

/* Example of a json request:

{"clientId":"1","routeId":"1"}
{"clientId":"2","routeId":"2"}
{"clientId":"3","routeId":"3"}

*/

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()

	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
