package nginx

import pb "go-agent/agent_proto"

type Server struct {
	pb.UnimplementedNginxServiceServer
}
