package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func connectNats() *nats.Conn {
	// User viper to import conf
	nc, err := nats.Connect("nats://nats-server:4222")
	if err != nil {
		log.Fatal("Could not connect to nats", err)
	} else {
		log.Println("Nats connected...")
	}
	return nc
}

/*
User service doesn't need to ask from transaction to solve some problems
*/

func performHealthCheck(nc *nats.Conn) {
	for {
		time.Sleep(2 * time.Second)
		msg, err := nc.Request("Ping.TransactionService", []byte("Ping"), 1000*time.Millisecond)
		if err != nil {
			log.Println("heathcheck failed: ", err)
		} else {
			log.Print("healthcheck passed: ", string(msg.Data))
		}
	}
}

func main() {
	nc := connectNats()

	go performHealthCheck(nc)

	// how will you retry if there is error ?
	// how to handle NATS errors ?
	fmt.Println("user service starting..")

	// add code to graceful shutdown
	time.Sleep(10000 * time.Second)
}
