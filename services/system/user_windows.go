//go:build windows

package system

import (
	"errors"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"go-agent/utils/windows"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func platformUserList(response *pb.UserListResponse) error {
	output, err := utils.RunCmd("net user")
	if err != nil {
		utils.LogError(err.Error())
		return err
	}
	users := strings.Split(string(output), "\n")[4:]

	for _, user := range users {
		user = strings.TrimSpace(user)
		saparator, _ := regexp.Compile(`\s+`)
		userSplit := saparator.Split(user, -1)
		if len(userSplit) < 3 {
			continue
		}
		user = userSplit[2]
		user = strings.Trim(user, "\r")
		if user == "" {
			continue
		}
		cmd := " useraccount where name=\"" + user + "\" get /value"
		utils.LogWarn(cmd)
		outputByte, err := windows.RunWmic(cmd)
		if err != nil {
			utils.LogError(err.Error())
			continue
		}
		//output, err = utils.GBKToUTF8(output)
		// convert to utf8 string to byte
		outputByte, err = utils.GBKToUTF8(outputByte)
		if err != nil {
			utils.LogError(err.Error())
			continue
		}
		lines := strings.Split(string(outputByte), "\n")
		userInfo := &pb.UserInfo{
			UserName: user,
		}

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			parts := strings.Split(line, "=")

			switch parts[0] {
			case "SID":
				userInfo.UserID = parts[1]
			case "PrimaryGroupId":
				userInfo.GroupID = parts[1]
			case "Description":
				userInfo.Comment = parts[1]
			case "HomeDirectory":

				//判断是否有home目录
				userInfo.HomeDir = parts[1]

			case "Name":
				userInfo.UserName = parts[1]
			}

		}
		isAdmin, err := windows.IsAdmin(userInfo.UserID)
		if err != nil {
			utils.LogError(err.Error())
			continue
		}
		if isAdmin {
			userInfo.UserType = "system"
		} else {
			userInfo.UserType = "user"
		}
		//get last login time
		lastLoginTime, err := getLastLoginTime(userInfo.UserName)
		if err == nil {
			userInfo.LastLoginTime = lastLoginTime
		}
		//判断当前用户是否在线
		if isOnline, err := isUserOnline(userInfo.UserName); err == nil {
			if isOnline {
				userInfo.Status = "Online"
			} else {
				userInfo.Status = "Offline"
			}
		}
		if userInfo.HomeDir == "" {
			// dir exist
			if _, err = os.Stat(filepath.Join("C:/Users", userInfo.UserName)); !os.IsNotExist(err) {
				userInfo.HomeDir = "C:/Users/" + userInfo.UserName
			}
		}
		response.List = append(response.List, userInfo)
	}

	return nil
}

func getLastLoginTime(username string) (string, error) {
	out, err := exec.Command("cmd", "/C", "net user "+username).Output()
	if err != nil {
		return "", err
	}
	out, _ = utils.GBKToUTF8(out)
	re := regexp.MustCompile(`Last logon\s*(.*)`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) < 2 {
		re = regexp.MustCompile(`上次登录\s*(.*)`)
		matches = re.FindStringSubmatch(string(out))
		if len(matches) < 2 {
			return "", errors.New("could not find last logon time")
		}
	}

	lastLogon := strings.TrimSpace(matches[1])

	return lastLogon, nil
}

func isUserOnline(username string) (bool, error) {
	//todo query command may not found in windows 10
	//qwinsta only works in windows server
	out, err := exec.Command("cmd", "/C", "query user").Output()
	if err != nil {
		return false, err
	}

	users := strings.Split(string(out), "\n")
	for _, user := range users {
		if strings.Contains(user, username) {
			return true, nil
		}
	}

	return false, nil
}
