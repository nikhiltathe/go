package main

import (
	"fmt"
	"time"

	"./kafka"
)

func main() {

	fmt.Println("OK")
	go kafka.StartKafka()
	fmt.Println("Kafka Started...")
	time.Sleep(5 * time.Minute)
}
