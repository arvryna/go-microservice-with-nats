package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/arvryna/betnomi/user-service/pb"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
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

// Grpc Stuffs
type UserManagerServer struct {
	pb.UnimplementedUserManagerServer
}

func (u *UserManagerServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Println("CreateUser gRPC request")
	return &pb.User{Name: in.Name, Email: in.Email, Token: uuid.New().String()}, nil
}

const PORT = ":9091"

func setupGRPCServer() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagerServer(s, &UserManagerServer{})

	log.Println("Trying to start gRPC server...")

	err = s.Serve(listener)
	if err != nil {
		log.Fatal("GPRC server failed to start", err)
	}
}

func main() {
	nc := connectNats()

	go performHealthCheck(nc)
	go setupGRPCServer()

	// how will you retry if there is error ?
	// how to handle NATS errors ?
	fmt.Println("user service starting..")

	// add code to graceful shutdown
	time.Sleep(10000 * time.Second)
}
