package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/arvryna/betnomi/user-service/db"
	"github.com/arvryna/betnomi/user-service/db/model"
	"github.com/arvryna/betnomi/user-service/pb"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
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

func (u *UserManagerServer) Login(ctx context.Context, in *empty.Empty) (*pb.LoginToken, error) {
	var user model.User
	user.Name = "user"
	user.Token = uuid.New().String()
	if res := DB.Create(&user); res.Error != nil {
		log.Println("User creation DB request failed: ", res.Error)
	}
	return &pb.LoginToken{Token: user.Token}, nil
}

// func getUserWithToken() model.User {

// }

func (u *UserManagerServer) Balance(ctx context.Context, in *pb.ExistingUser) (*pb.UserBalance, error) {
	var user model.User
	fmt.Println(in.Token)
	DB.Where("token = ?", in.Token).Find(&user)
	return &pb.UserBalance{Balance: user.Balance}, nil
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
	DB = db.Init()

	nc := connectNats()

	go performHealthCheck(nc)

	go setupGRPCServer()

	// how will you retry if there is error ?
	// how to handle NATS errors ?
	fmt.Println("user service starting..")

	// add code to graceful shutdown
	time.Sleep(10000 * time.Second)
}
