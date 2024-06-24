//go:build windows

package get_apache_info

import pb "go-agent/agent_proto"

func PlatformGetApacheInfo() (*pb.GetApacheInfoResponse, error) {
	var response pb.GetApacheInfoResponse
	var err error
	response.Message = "unimplemented"
	return &response, err
}
