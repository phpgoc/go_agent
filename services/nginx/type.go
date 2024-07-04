package nginx

import (
	pb "go-agent/agent_proto"
	"strings"
)

type Server struct {
	pb.UnimplementedNginxServiceServer
}

type InfoResponseWrapper struct {
	*pb.GetNginxInfoResponse
}

func (w InfoResponseWrapper) String() string {
	var r = w.GetNginxInfoResponse
	var sb strings.Builder
	sb.WriteString("GetNginxInfoResponse{")
	sb.WriteString("Message: " + r.Message + ", ")
	sb.WriteString("NginxInstances: [")
	for i, instance := range r.NginxInstances {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("{ConfigPath: " + instance.ConfigPath + ", ")
		sb.WriteString("ErrorLog: {FilePath: " + instance.ErrorLog.FilePath + ", Size: " + instance.ErrorLog.Size + ", ModifyTime: " + instance.ErrorLog.ModifyTime + "}, ")
		sb.WriteString("AccessLog: {FilePath: " + instance.AccessLog.FilePath + ", Size: " + instance.AccessLog.Size + ", ModifyTime: " + instance.AccessLog.ModifyTime + "}}")
	}
	sb.WriteString("]}")
	return sb.String()
}
