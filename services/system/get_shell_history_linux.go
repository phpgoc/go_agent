package system

import (
	"fmt"
	pb "go-agent/agent_proto"
	"go-agent/utils"

	"regexp"
	"strings"
	"unicode/utf8"
)

func platformGetShellHistory(req *pb.GetShellHistoryRequest) (*pb.GetShellHistoryResponse, error) {
	var res pb.GetShellHistoryResponse

	var userList pb.UserListResponse
	//有错userList会是空的，没关系
	_ = platformUserList(&userList)

	if req.UserName != "" {
		for _, user := range userList.List {
			if user.UserName == req.UserName {
				history, err := getCommandHistoryByUser(req.UserName, user.HomeDir)
				if err != nil {
					res.Message = err.Error()
					return &res, err
				} else {
					res.ListByUserName = append(res.ListByUserName, history)
					return &res, nil
				}
			}
		}
	} else {
		for _, user := range userList.List {
			history, err := getCommandHistoryByUser(user.UserName, user.HomeDir)
			if err != nil {
				res.Message = err.Error()
				return &res, err
			} else {
				res.ListByUserName = append(res.ListByUserName, history)
			}
		}
	}

	return &res, nil
}

func getCommandHistoryByUser(userName, homeDir string) (*pb.UserHistory, error) {
	res := pb.UserHistory{}

	filesSource := utils.GetFirstAndLogError(
		func() (string, error) {
			return utils.RunCmd(fmt.Sprintf("find %v  -name '.*_history'", homeDir))
		})
	if filesSource == "" {
		//root可能是开发环境运行，没权利访问其他用户的目录
		utils.LogWarn(fmt.Sprintf("user %v home not found", userName))
	} else {
		files := strings.Split(filesSource, "\n")
		re, _ := regexp.Compile(`/\.([^_]+)_history$`)
		for _, file := range files {
			if file == "" {
				continue
			}
			match := re.FindStringSubmatch(file)
			if len(match) > 1 {
				shellName := match[1]
				commandsSource, _ := utils.RunCmd(fmt.Sprintf("cat %v", file))
				commandsString := strings.Split(commandsSource, "\n")
				newCommandsString := make([]string, 0)
				for _, command := range commandsString {
					// assert command is valid utf-8
					commandBytes := []byte(command)
					if !utf8.Valid(commandBytes) {
						utils.LogWarn(fmt.Sprintf("invalid utf-8 command: %v", command))
						continue
					}
					newCommandsString = append(newCommandsString, command)
				}
				res.ListByShellName = append(res.ListByShellName, &pb.ShellHistory{
					ShellName:     shellName,
					ListByCommand: newCommandsString,
				})
			}
		}
	}

	res.UserName = userName
	return &res, nil
}
