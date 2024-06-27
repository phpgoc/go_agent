package main

import (
	"flag"
	"fmt"
	"go-agent/services/get_apache_info"
	"go-agent/services/get_sys_info"
	"go-agent/services/helloworld"
	"go-agent/services/network"
	"go-agent/services/user_list"
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
		log.Printf(err.Error())
		return
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		utils.LogError(fmt.Sprintf("failed to listen: %v", err))
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &helloworld.Server{})
	pb.RegisterGetApacheInfoServer(s, &get_apache_info.Server{})
	pb.RegisterGetSysInfoServer(s, &get_sys_info.Server{})
	pb.RegisterGetUserListServer(s, &user_list.Server{})
	pb.RegisterNetworkServer(s, &network.Server{})
	utils.LogInfo(fmt.Sprintf("server listening at %v", lis.Addr()))
	if err := s.Serve(lis); err != nil {
		utils.LogError(fmt.Sprintf("failed to serve: %v", err))
	}
}
