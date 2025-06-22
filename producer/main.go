package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Mood struct {
    Name      string `json:"name"`
    Mood      string `json:"mood"`
    Timestamp string `json:"timestamp"`
}

var producer *kafka.Producer

func moodHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Received /mood request")

    var m Mood
    if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
        log.Printf("Decode error: %v", err)
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    m.Timestamp = r.Header.Get("X-Timestamp")

    moodJSON, err := json.Marshal(m)
    if err != nil {
        log.Printf("JSON marshal error: %v", err)
        http.Error(w, "Failed to process mood", http.StatusInternalServerError)
        return
    }

    topic := "mood"
    err = producer.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:          moodJSON,
    }, nil)
    if err != nil {
        log.Printf("Kafka produce error: %v", err)
        http.Error(w, "Failed to produce message", http.StatusInternalServerError)
        return
    }

    log.Printf("Mood successfully produced: %s", string(moodJSON))
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte("Mood received"))
}

func main() {
    var err error
    producer, err = kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "kafka:9092",
    })
    if err != nil {
        log.Fatalf("Failed to create producer: %s", err)
    }
    defer producer.Close()

    go func() {
        for e := range producer.Events() {
            switch ev := e.(type) {
            case *kafka.Message:
                if ev.TopicPartition.Error != nil {
                    log.Printf("Delivery failed: %v", ev.TopicPartition.Error)
                } else {
                    log.Printf("Message delivered to %v", ev.TopicPartition)
                }
            }
        }
    }()

    http.HandleFunc("/mood", moodHandler)
    log.Println("Starting producer on :61810...")
    if err := http.ListenAndServe(":61810", nil); err != nil {
        log.Fatalf("HTTP server error: %v", err)
    }
}
