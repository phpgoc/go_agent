package system

import (
	"context"
	pb "go-agent/agent_proto"
)

func (s *Server) GetSystemServices(_ context.Context, _ *pb.GetSystemServicesRequest) (*pb.GetSystemServicesResponse, error) {
	response := pb.GetSystemServicesResponse{}

	return platformGetSystemServices(&response)
}
