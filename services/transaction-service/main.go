package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/arvryna/betnomi/transaction-service/pb"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
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
	balance := int64(101)
	return &pb.TransactionResponse{Balance: balance}, nil
}
func (t *TransactionManagerServer) TransactionDown(ctx context.Context, in *pb.NewTransaction) (*pb.TransactionResponse, error) {
	balance := int64(101)
	return &pb.TransactionResponse{Balance: balance}, nil
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
	nc := connectNats()

	nc.Subscribe("Ping.TransactionService", func(m *nats.Msg) {
		if string(m.Data) == "Ping" {
			m.Respond([]byte("Pong"))
		}
	})

	go setupGRPCServer()

	fmt.Println("transaction service started...")

	time.Sleep(1000 * time.Second)

	//handling graceful shutdown
}
