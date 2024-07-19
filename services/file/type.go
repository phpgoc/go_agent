package file

import pb "go-agent/agent_proto"

type Server struct {
	pb.UnimplementedFileServiceServer
}
