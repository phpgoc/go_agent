package database

import (
	"context"
	pb "go-agent/agent_proto/database"
)

var ConnectionInfoSet = make(map[ConnectionInfoForKey]bool)

func (s *Server) MysqlDump(_ context.Context, request *pb.MysqlDumpRequest) (*pb.MysqlDumpResponse, error) {

	return nil, nil
}

func canConnect(info *ConnectionInfoForKey) bool {
	return false
}

func dump(ip, port, username, password string) (file string) {
	return ""
}
