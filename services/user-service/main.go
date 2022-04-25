package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
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

func getUserWithId(id int) model.User {
	var user model.User
	DB.Where("id = ?", id).Find(&user)
	return user
}

func getUserWithToken(token string) model.User {
	var user model.User
	DB.Where("token = ?", token).Find(&user)
	return user
}

func (u *UserManagerServer) Balance(ctx context.Context, in *pb.ExistingUser) (*pb.UserBalance, error) {
	balance := getUserWithToken(in.Token).Balance
	return &pb.UserBalance{Balance: balance}, nil
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

// DTO
type TransactionDto struct {
	Id                int `gorm:"primaryKey"`
	TransactionAmount int64
	IsUp              bool
	UserId            int
	Before            int64
	After             int64
	CreatedAt         time.Time
}

func main() {
	DB = db.Init()
	nc := connectNats()

	// gets token, sends userId
	nc.Subscribe("userservice.getuserid", func(m *nats.Msg) {
		id := getUserWithToken(string(m.Data)).Id
		m.Respond([]byte(strconv.Itoa(id)))
	})

	nc.Subscribe("userservice.getuserbalance", func(m *nats.Msg) {
		balance := getUserWithToken(string(m.Data)).Balance
		m.Respond([]byte(strconv.Itoa(int(balance))))
	})

	nc.Subscribe("userservice.updatebalance", func(m *nats.Msg) {
		var transaction TransactionDto
		json.Unmarshal(m.Data, &transaction)
		user := getUserWithId(transaction.UserId)
		user.Balance = transaction.After
		DB.Save(&user)
	})

	go performHealthCheck(nc)

	go setupGRPCServer()

	// how will you retry if there is error ?
	// how to handle NATS errors ?
	fmt.Println("user service starting..")

	// add code to graceful shutdown
	time.Sleep(10000 * time.Second)
}
