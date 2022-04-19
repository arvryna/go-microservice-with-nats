package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func connectNats() *nats.Conn {
	nc, err := nats.Connect("nats://nats-server:4222")
	if err != nil {
		//TODO: add proper logging to identify service names
		log.Fatal("Could not connect to nats")
	} else {
		log.Println("Nats connected...")
	}
	return nc
}

func main() {
	nc := connectNats()
	nc.Subscribe("Ping.TransactionService", func(m *nats.Msg) {
		if string(m.Data) == "Ping" {
			m.Respond([]byte("Pong"))
		}
	})
	fmt.Println("transaction service started...")

	time.Sleep(1000 * time.Second)

	//handling graceful shutdown
}
