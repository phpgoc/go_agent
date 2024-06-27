package user_list

import (
	"context"
	pb "go-agent/agent_proto"
	"go-agent/utils"
)

type Server struct {
	pb.UnimplementedGetUserListServer
}

func (s *Server) GetUserList(_ context.Context, _ *pb.UserListRequest) (*pb.UserListResponse, error) {
	utils.LogInfo("call GetUserList")
	var response pb.UserListResponse
	err := platformUserList(&response)
	if err != nil {
		utils.LogError(err.Error())
		return nil, err
	}
	return &response, nil
}
