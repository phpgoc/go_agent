package main

import (
	"flag"
	"fmt"
	"go-agent/services/get_apache_info"
	"go-agent/services/helloworld"
	"go-agent/utils"
	"log"
	"net"

	pb "go-agent/agent_proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()

	var err = utils.Init()
	if err != nil {
		err := utils.LogError(err.Error())
		if err != nil {
			return
		}
		return
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &helloworld.GreeterServer{})
	pb.RegisterGetApacheInfoServer(s, &get_apache_info.GetApacheInfoServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
