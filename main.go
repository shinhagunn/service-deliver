package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/shinhagunn/service-deliver/delivery"
	"google.golang.org/grpc"
)

type Service interface {
	Process()
}

type server struct {
	pb.UnimplementedDeliverServer
}

func (s *server) Deliver(ctx context.Context, in *pb.DeliverRequest) (*pb.DeliverResponse, error) {
	log.Printf("Da nhan duoc don hang cua: %v", in.GetName())
	time.Sleep(time.Minute)

	return &pb.DeliverResponse{Status: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterDeliverServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
