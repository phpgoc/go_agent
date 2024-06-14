package main

import (
	"flag"
	"fmt"
	"go-agent/command/helloworld"
	"go-agent/utils"
	"log"
	"net"

	pb "go-agent/proto_go"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()

	var err = utils.Init()
	if err != nil {
		utils.LogError(err.Error())
		return
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &helloworld.GreeterServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
