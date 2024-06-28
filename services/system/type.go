package system

import (
	pb "go-agent/agent_proto"
)

type Server struct {
	pb.UnimplementedSystemServiceServer
}
