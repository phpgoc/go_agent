package system

import (
	"fmt"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"strings"
)

func platformGetSystemServices(response *pb.GetSystemServicesResponse) (*pb.GetSystemServicesResponse, error) {
	//cmd := "systemctl list-units --type=service --no-legend --no-pager"
	cmd := "systemctl list-units --no-legend --no-pager" //这个命令会列出所有的服务
	output, err := utils.RunCmd(cmd)
	if err != nil {
		return utils.SetResponseErrorAndLogMessageGeneric(response, fmt.Sprintf("failed to execute command: %v", err), pb.ResponseCode_UNKNOWN_SERVER_ERROR)
	}

	for _, line := range strings.Split(output, "\n") {
		line = strings.TrimSpace(line)
		fields := utils.SplitFields(line)
		if len(fields) < 5 {
			continue
		}
		service := pb.SystemServiceInfo{
			//前边有空格所以
			Name:        fields[0],
			Description: strings.Join(fields[4:], " "),
			State:       fields[3],
		}
		response.List = append(response.List, &service)
	}
	utils.LogInfo(fmt.Sprintf("GetSystemServices response len: %d", len(response.List)))
	return response, nil
}
