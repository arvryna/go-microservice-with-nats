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
		log.Fatal("Could not connect to nats")
		panic(err)
	} else {
		log.Println("Nats connected...")
	}
	return nc
}

/*
User service doesn't need to ask from transaction to solve some problems
*/

func main() {
	nc := connectNats()

	// how will you retry if there is error ?
	// how to handle NATS errors ?
	msg, err := nc.Request("Ping.TransactionService", nil, 10*time.Millisecond)

	if err != nil {
		log.Fatal("Transervice Ping failed.. ", err)
	} else {
		log.Println("Ping reponse", string(msg.Data))
	}

	fmt.Println("user service starting..")

	// add code to graceful shutdown
}
