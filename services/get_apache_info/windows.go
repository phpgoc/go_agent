//go:build windows

package get_apache_info

import (
	pb "go-agent/agent_proto"
	"go-agent/utils"
)

func PlatformGetApacheInfo() (*pb.GetApacheInfoResponse, error) {
	var response pb.GetApacheInfoResponse
	var err error
	response.Message = "unimplemented"
	utils.LogInfo(response.String())
	return &response, err
}
