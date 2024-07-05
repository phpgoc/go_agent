package main

import (
	"flag"
	"fmt"
	"go-agent/agent_runtime"
	"go-agent/services/apache"
	"go-agent/services/file"
	"go-agent/services/helloworld"
	"go-agent/services/network"
	"go-agent/services/nginx"
	"go-agent/services/system"
	"go-agent/utils"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"

	pb "go-agent/agent_proto"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	var err = utils.Init()
	if err != nil {
		//这里不使用utils里LogError是因Log Writer可能初始化失败
		log.Printf(err.Error())
		os.Exit(1)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *agent_runtime.Port))
	if err != nil {
		utils.LogError(fmt.Sprintf("failed to listen: %v", err))
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &helloworld.Server{})
	pb.RegisterApacheServiceServer(s, &apache.Server{})
	pb.RegisterSystemServiceServer(s, &system.Server{})
	pb.RegisterNetworkServiceServer(s, &network.Server{})
	pb.RegisterFileServiceServer(s, &file.Server{})
	pb.RegisterNginxServiceServer(s, &nginx.Server{})

	//grpcurl --plaintext localhost:50051 list
	//增加反射服务,客户端可以通过反射服务发现服务
	reflection.Register(s)

	utils.LogInfo(fmt.Sprintf("server listening at %v", lis.Addr()))
	if err := s.Serve(lis); err != nil {
		utils.LogError(fmt.Sprintf("failed to serve: %v", err))
	}
}
