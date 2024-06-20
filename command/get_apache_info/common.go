package get_apache_info

import (
	"context"
	pb "go-agent/agent_proto"
)

type GetApacheInfoServer struct {
	pb.UnimplementedGetApacheInfoServer
}

func (s *GetApacheInfoServer) GetApacheInfo(_ context.Context, in *pb.GetApacheInfoRequest) (*pb.GetApacheInfoResponse, error) {
	return PlatformGetApacheInfo()
}

func recursiveInsertData(file string, response pb.GetApacheInfoResponse) {

}
