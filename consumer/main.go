// package main

// import (
// 	"fmt"
// 	"time"

// 	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
// )

// func main() {

// 	c, err := kafka.NewConsumer(&kafka.ConfigMap{
// 		"bootstrap.servers": "localhost:9093",
// 		"group.id":          "myGroup",
// 		"auto.offset.reset": "earliest",
// 	})

// 	if err != nil {
// 		panic(err)
// 	}

// 	err = c.SubscribeTopics([]string{"topic2", "^aRegex.*[Tt]opic"}, nil)

// 	if err != nil {
// 		panic(err)
// 	}

// 	// A signal handler or similar could be used to set this to false to break the loop.
// 	run := true

// 	for run {
// 		msg, err := c.ReadMessage(time.Second)
// 		if err == nil {
// 			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
// 		} else if !err.(kafka.Error).IsTimeout() {
// 			// The client will automatically try to recover from all errors.
// 			// Timeout is not considered an error because it is raised by
// 			// ReadMessage in absence of messages.
// 			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
// 		}
// 	}

// 	c.Close()
// }

package main

import (
    "fmt"
    "time"

    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "kafka:9092",
        "group.id":          "mood-consumer-group",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        panic(err)
    }

    c.SubscribeTopics([]string{"mood"}, nil)
    fmt.Println("Mood consumer started...")

    for {
        msg, err := c.ReadMessage(time.Second)
        if err == nil {
            fmt.Printf("Received: %s\n", msg.Value)
        }
    }
}
