package get_apache_info

import (
	"context"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"strconv"
	"strings"
)

type GetApacheInfoServer struct {
	pb.UnimplementedGetApacheInfoServer
}

func (s *GetApacheInfoServer) GetApacheInfo(_ context.Context, in *pb.GetApacheInfoRequest) (*pb.GetApacheInfoResponse, error) {
	return PlatformGetApacheInfo()
}

func recursiveInsertData(fileName, rootPath string, response pb.GetApacheInfoResponse) (err error) {
	var fileContent string
	if fileContent, err = utils.ReadFile(fileName); err != nil {
		response.Message = "can't read file " + fileName
		return err
	}
	//find all include file_name
	var includeOptions []string

	//readline'
	var site pb.SiteInfo
	for _, line := range strings.Split(fileContent, "\n") {
		//if start with # ,continue
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		if strings.Contains(line, "Include") {
			//if have no two element, continue
			if option := utils.SplitStringAndGetIndexSafe(line, " ", 1); option == "" {
				includeOptions = append(includeOptions, strings.Trim(option, "\""))
			}
		}
		if strings.Contains(line, "<VirtualHost") {
			site = pb.SiteInfo{}
		}
		if strings.Contains(line, "</VirtualHost>") {
			response.List = append(response.List, &site)
		}
		if strings.Contains(line, "VirtualHost") {
			listenString := strings.Trim(utils.SplitStringAndGetIndexSafe(line, " ", 1), "\"")
			parseInt, _ := strconv.ParseInt(listenString, 10, 32)
			site.Listen = int32(parseInt)
		}
		if strings.Contains(line, "ServerName") {
			site.ServerName = utils.SplitStringAndGetIndexSafe(line, " ", 1)
		}
		if strings.Contains(line, "DocumentRoot") {
			site.DocumentRoot = utils.SplitStringAndGetIndexSafe(line, " ", 1)
		}
		if strings.Contains(line, "ServerAlias") {
			site.ServerAlias = utils.SplitStringAndGetIndexSafe(line, " ", 1)
		}
		if strings.Contains(line, "ErrorLog") {
			//如果有错，内容会全空
			size, accessTime, modifyTime := utils.ExtractFileStat(utils.SplitStringAndGetIndexSafe(line, " ", 1))
			var log = pb.Log{
				Type:         "错误",
				Size:         size,
				AccessTime:   accessTime,
				ModifiedTime: modifyTime,
			}
			site.Logs = append(site.Logs, &log)
		}
		if strings.Contains(line, "CustomLog") {
			size, accessTime, modifyTime := utils.ExtractFileStat(utils.SplitStringAndGetIndexSafe(line, " ", 1))
			var log = pb.Log{
				Type:         "访问",
				Size:         size,
				AccessTime:   accessTime,
				ModifiedTime: modifyTime,
			}
			site.Logs = append(site.Logs, &log)
		}
	}
	// if includeOptions is has * ,find all match file

	for _, option := range includeOptions {
		if strings.Contains(option, "*") {

			files, err := utils.FindMatchedFiles(option)
			if err != nil {
				return err
			}
			for _, file := range files {
				err := recursiveInsertData(file, rootPath, response)
				if err != nil {
					return err
				}
			}

		} else {
			err := recursiveInsertData(option, rootPath, response)
			if err != nil {
				return err
			}
		}
	}
	return
}
