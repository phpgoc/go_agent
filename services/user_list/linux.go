//go:build linux

package user_list

import (
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"strconv"
	"strings"
)

func platformUserList(response *pb.UserListResponse) error {
	groupSource, err := utils.RunCmd("cat /etc/group")
	if err != nil {
		utils.LogError(err.Error())
		return err
	}
	// key group id，value group name
	groups := make(map[string]string)
	for _, line := range strings.Split(groupSource, "\n") {
		if line == "" {
			continue
		}
		groupInfo := strings.Split(line, ":")
		groups[groupInfo[2]] = groupInfo[0]
	}
	usersSource, err := utils.RunCmd("cat /etc/passwd")
	if err != nil {
		utils.LogError(err.Error())
		return err
	}

	//这段代码不用safe split 也不处理error，因为这个是系统文件，不会有恶意输入
	for _, line := range strings.Split(usersSource, "\n") {
		//skip nologin
		if strings.Contains(line, "/sbin/nologin") {
			continue
		}
		arr := strings.Split(line, ":")
		if len(arr) < 7 {
			continue
		}

		userInfo := &pb.UserInfo{
			UserName:  arr[0],
			UserID:    arr[2],
			GroupID:   arr[3],
			GroupName: groups[arr[3]],
			Comment:   arr[4],
			HomeDir:   arr[5],
			Shell:     arr[6],
		}
		userId2Int, _ := strconv.Atoi(arr[2])
		//convert to int
		if userId2Int < 1000 {
			userInfo.UserType = "system"
		} else {
			userInfo.UserType = "user"
		}
		lastLoginSource, _ := utils.RunCmd("lastlog -u " + arr[0])
		//获取第二行的后半段
		lastLoginStr := strings.Split(strings.Split(lastLoginSource, "\n")[1], arr[0])[1]
		userInfo.LastLoginTime = strings.TrimSpace(lastLoginStr)
		response.List = append(response.List, userInfo)
	}
	return nil
}
