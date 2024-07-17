package system

import (
	"context"
	"fmt"
	pb "go-agent/agent_proto"
	"go-agent/utils"
)

func (s *Server) GetShellHistory(_ context.Context, req *pb.GetShellHistoryRequest) (*pb.GetShellHistoryResponse, error) {
	utils.LogInfo(fmt.Sprintf("GetShellHistory: %v", req))
	responsePtr, err := platformGetShellHistory(req)
	if err != nil {
		return utils.SetResponseErrorAndLogMessageGeneric(responsePtr, err.Error(), pb.ResponseCode_UNKNOWN_SERVER_ERROR)
	}
	var count = 0
	for _, userHistory := range responsePtr.ListByUserName {
		for _, shellHistory := range userHistory.ListByShellName {
			count += len(shellHistory.ListByCommand)
		}
	}
	utils.LogInfo(fmt.Sprintf("GetShellHistory history count: %d", count))
	return responsePtr, nil
}
