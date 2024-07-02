package system

import (
	"fmt"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"strings"
	"unicode/utf8"
)

func platformGetShellHistory(req *pb.GetShellHistoryRequest) (*pb.GetShellHistoryResponse, error) {
	var res pb.GetShellHistoryResponse

	var userList pb.UserListResponse
	//有错userList会是空的，没关系
	_ = platformUserList(&userList)

	if req.UserName != "" {
		filePath := fmt.Sprintf("C:/Users/%s/AppData/Roaming/Microsoft/Windows/PowerShell/PSReadLine/ConsoleHost_history.txt",
			req.UserName)
		if !utils.FileExists(filePath) {
			res.Message = fmt.Sprintf("user %s history file not found", req.UserName)
			return &res, nil
		}
		addCommandHistoryByFile(filePath, req.UserName, &res)
	} else {
		for _, user := range userList.List {
			filePath := fmt.Sprintf("C:/Users/%s/AppData/Roaming/Microsoft/Windows/PowerShell/PSReadLine/ConsoleHost_history.txt",
				user.UserName)
			if !utils.FileExists(filePath) {
				utils.LogWarn(fmt.Sprintf("user %s history file not found", user.UserName))
				continue
			}
			addCommandHistoryByFile(filePath, user.UserName, &res)

		}

	}

	return &res, nil
}

// filePath肯定是存在的
func addCommandHistoryByFile(filePath, userName string, res *pb.GetShellHistoryResponse) {
	fileContent, _ := utils.ReadFile(filePath)
	UserHistory := pb.UserHistory{
		UserName: userName,
		ListByShellName: []*pb.ShellHistory{
			{
				ShellName: "powershell",
			},
		},
	}
	for _, line := range strings.Split(fileContent, "\n") {
		if line == "" {
			continue
		}
		lineByte := []byte(line)
		if !utf8.Valid(lineByte) {
			utils.LogWarn(fmt.Sprintf("invalid utf8: %v", line))
			continue
		}
		UserHistory.ListByShellName[0].ListByCommand = append(UserHistory.ListByShellName[0].ListByCommand, line)
	}

	res.ListByUserName = append(res.ListByUserName, &UserHistory)
}
