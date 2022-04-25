package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/arvryna/betnomi/transaction-service/db"
	"github.com/arvryna/betnomi/transaction-service/db/model"
	"github.com/arvryna/betnomi/transaction-service/pb"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	NC *nats.Conn
	DB *gorm.DB
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

// Grpc Stuffs
type TransactionManagerServer struct {
	pb.UnimplementedTransactionManagerServer
}

func (t *TransactionManagerServer) TransactionUp(ctx context.Context, in *pb.NewTransaction) (*pb.TransactionResponse, error) {
	// Get user ID
	msg, _ := NC.Request("userservice.getuserid", []byte(in.Token), 1000*time.Millisecond)
	id, _ := strconv.Atoi(string(msg.Data))

	// Get user Balance
	msg, _ = NC.Request("userservice.getuserbalance", []byte(in.Token), 1000*time.Millisecond)
	balance, _ := strconv.Atoi(string(msg.Data))

	var transaction model.Transaction
	transaction.TransactionAmount = in.Value
	transaction.UserId = id
	transaction.Before = int64(balance)
	transaction.After = int64(balance) + in.Value // Take existing balance and add new value to it
	transaction.IsUp = true

	b, _ := json.Marshal(transaction)
	// Get user Balance
	NC.Request("userservice.updatebalance", b, 1000*time.Millisecond)

	if res := DB.Create(&transaction); res.Error != nil {
		log.Println("Transaction creation DB request failed: ", res.Error)
	}

	return &pb.TransactionResponse{Balance: int64(transaction.After)}, nil
}

func (t *TransactionManagerServer) TransactionDown(ctx context.Context, in *pb.NewTransaction) (*pb.TransactionResponse, error) {

	msg, _ := NC.Request("userservice.getuserid", []byte(in.Token), 1000*time.Millisecond)
	id, _ := strconv.Atoi(string(msg.Data))

	// Get user Balance
	msg, _ = NC.Request("userservice.getuserbalance", []byte(in.Token), 1000*time.Millisecond)
	balance, _ := strconv.Atoi(string(msg.Data))

	var transaction model.Transaction
	transaction.TransactionAmount = in.Value
	transaction.UserId = id
	transaction.Before = int64(balance)
	transaction.After = int64(balance) - in.Value // Take existing balance and subtract new value to it
	transaction.IsUp = false

	b, _ := json.Marshal(transaction)
	// Get user Balance
	NC.Request("userservice.updatebalance", b, 1000*time.Millisecond)

	if res := DB.Create(&transaction); res.Error != nil {
		log.Println("Transaction creation DB request failed: ", res.Error)
	}

	return &pb.TransactionResponse{Balance: int64(transaction.After)}, nil
}

const PORT = ":9092"

func setupGRPCServer() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterTransactionManagerServer(s, &TransactionManagerServer{})

	log.Println("Trying to start gRPC server...")

	err = s.Serve(listener)
	if err != nil {
		log.Fatal("GPRC server failed to start", err)
	}
}

func main() {
	DB = db.Init()
	NC = connectNats()

	NC.Subscribe("Ping.TransactionService", func(m *nats.Msg) {
		if string(m.Data) == "Ping" {
			m.Respond([]byte("Pong"))
		}
	})

	go setupGRPCServer()

	fmt.Println("transaction service started...")

	time.Sleep(1000 * time.Second)

	//handling graceful shutdown
}
