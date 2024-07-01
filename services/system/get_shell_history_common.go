package system

import (
	"context"
	"fmt"
	pb "go-agent/agent_proto"
	"go-agent/utils"
)

func (s *Server) GetShellHistory(_ context.Context, req *pb.GetShellHistoryRequest) (*pb.GetShellHistoryResponse, error) {
	utils.LogInfo(fmt.Sprintf("GetShellHistory: %v", req))
	res, err := platformGetShellHistory(req)
	if err != nil {
		res.Message = err.Error()
		utils.LogError(res.Message)
		return res, nil
	}
	var count = 0
	for _, userHistory := range res.ListByUserName {
		for _, shellHistory := range userHistory.ListByShellName {
			count += len(shellHistory.ListByCommand)
		}
	}
	utils.LogInfo(fmt.Sprintf("history count: %v", count))
	return res, nil
}
