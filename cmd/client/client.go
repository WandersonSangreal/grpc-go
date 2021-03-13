package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/codeedu/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	// AddUser(client)
	// AddUserVerbose(client)
	AddUsers(client)

}

func AddUser(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Wanderson Sangreal",
		Email: "sangreal@wanderson.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	log.Println(res)

}

func AddUserVerbose(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Wanderson Sangreal",
		Email: "sangreal@wanderson.com",
	}

	res, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for {

		stream, err := res.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}

		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())

	}

}

func AddUsers(client pb.UserServiceClient) {

	reqs := []*pb.User{
		{
			Id:    "0",
			Name:  "Wanderson Sangreal",
			Email: "sangreal@wanderson.com",
		},
		{
			Id:    "1",
			Name:  "Wanderson Sangreal 02",
			Email: "sangreal01@wanderson.com",
		},
		{
			Id:    "2",
			Name:  "Wanderson Sangreal 03",
			Email: "sangreal02@wanderson.com",
		},
		{
			Id:    "3",
			Name:  "Wanderson Sangreal 04",
			Email: "sangreal03@wanderson.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {

		stream.Send(req)
		time.Sleep(time.Second * 3)

	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)

}
