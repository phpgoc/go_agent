package database

import (
	"fmt"
	pb "go-agent/agent_proto/database"
)

type Server struct {
	pb.UnimplementedDatabaseServiceServer
}

type ConnectionInfoForKey struct {
	Username        string
	Password        string
	Host            string
	Port            uint32
	SkipGrantTables bool
}

type mysqlDumpRequestWrapper struct {
	*pb.MysqlDumpRequest
}

func (w mysqlDumpRequestWrapper) String() string {
	var r = w.MysqlDumpRequest
	var res = fmt.Sprintf("MysqlDumpRequest{SkipGrantTables: %v, Force: %v", r.SkipGrantTables, r.Force)
	if r.ConnectionInfo != nil {
		res += fmt.Sprintf(", ConnectionInfo: {Username: %s, Password: %s, Ip: %s, Port: %d}", r.ConnectionInfo.Username, r.ConnectionInfo.Password, r.ConnectionInfo.Host, r.ConnectionInfo.Port)
	}
	res += "}"
	return res
}
