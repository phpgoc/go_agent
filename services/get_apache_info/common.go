package get_apache_info

import (
	"context"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"path/filepath"
	"strconv"
	"strings"
)

type Server struct {
	pb.UnimplementedGetApacheInfoServer
}

func (s *Server) GetApacheInfo(_ context.Context, _ *pb.GetApacheInfoRequest) (*pb.GetApacheInfoResponse, error) {
	utils.LogInfo("called ApacheInfo")
	var response pb.GetApacheInfoResponse
	var err error
	var apacheV string

	//大V可能会返回错误，即使有apache，报错还不不影响apache正常运行，所用用小v判断是否存在，如果存在则用大V
	if _, err = utils.RunCmd("apache2 -v"); err == nil {
		apacheV, err = utils.RunCmd("apache2 -V")
	} else if apacheV, err = utils.RunCmd("apache -v"); err == nil {
		apacheV, _ = utils.RunCmd("apache -V")
	} else if apacheV, err = utils.RunCmd("httpd -v"); err == nil {
		apacheV, _ = utils.RunCmd("httpd -V")
	} else {
		response.Message = "can't find apache"
		return &response, err
	}
	apacheCmd := utils.FindCommandFromPathAndProcessByMatchStringArray([]string{"apache2", "httpd", "apache"})
	if apacheCmd == "" {
		response.Message = "can't find apache"
		utils.LogError(response.Message)
		return &response, err
	}

	apacheV, err = utils.RunCmd(apacheCmd + " -V")
	if err != nil {
		response.Message = "apache compiler error or loss dependency"
		utils.LogError(response.Message)
		//不返回
	}
	var httpRoot, serverConfig string
	for _, line := range strings.Split(apacheV, "\n") {
		if strings.Contains(line, "HTTPD_ROOT") {
			if httpRoot = utils.SplitStringAndGetIndexSafely(line, "=", 1); httpRoot != "" {
				httpRoot = strings.Trim(httpRoot, "\"")

			} else {
				break
			}
		}
		if strings.Contains(line, "SERVER_CONFIG_FILE") {
			if serverConfig = utils.SplitStringAndGetIndexSafely(line, "=", 1); serverConfig != "" {
				serverConfig = strings.Trim(serverConfig, "\"")
				//set empty string to split
			} else {
				break
			}
		}
	}
	if httpRoot == "" || serverConfig == "" {
		response.Message = "can't find HTTPD_ROOT or SERVER_CONFIG_FILE"
		return &response, err
	}
	file := filepath.Join(httpRoot, serverConfig)
	//env dict ,this file name by guess
	envContent, _ := utils.ReadFile(filepath.Join(httpRoot, "envvars"))
	envMap := utils.InterpretSourceExportToGoMap(envContent, map[string]string{})
	err = recursiveInsertData(file, httpRoot, &response, envMap)
	if err != nil {
		utils.LogError(err.Error())
		return nil, err
	}
	utils.LogInfo(response.String())
	return &response, nil
}

func recursiveInsertData(fileName, rootPath string, response *pb.GetApacheInfoResponse, envMap map[string]string) (err error) {
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
		if strings.Contains(line, "IncludeOptional") {
			//if have no two element, continue
			if option := utils.SplitStringAndGetIndexSafely(line, " ", 1); option != "" {
				//避免加载到loadmodule
				if strings.HasSuffix(option, "load") {
					continue
				}
				includeOptions = append(includeOptions, strings.Trim(option, "\""))
			}
			continue
		}
		if strings.Contains(line, "<VirtualHost") {
			site = pb.SiteInfo{}
			continue
		}
		if strings.Contains(line, "</VirtualHost>") {
			response.List = append(response.List, &site)
			continue
		}
		if strings.Contains(line, "VirtualHost") {
			listenString := strings.Trim(utils.SplitStringAndGetIndexSafely(line, " ", 1), "\"")
			parseInt, _ := strconv.ParseInt(listenString, 10, 32)
			site.Listen = int32(parseInt)

		}
		if strings.Contains(line, "ServerName") {
			site.ServerName = utils.SplitStringAndGetIndexSafely(line, " ", 1)
		}
		if strings.Contains(line, "DocumentRoot") {
			site.DocumentRoot = utils.SplitStringAndGetIndexSafely(line, " ", 1)
		}
		if strings.Contains(line, "ServerAlias") {
			site.ServerAlias = utils.SplitStringAndGetIndexSafely(line, " ", 1)
		}
		if strings.Contains(line, "ErrorLog") {
			//如果有错，内容会全空
			size, accessTime, modifyTime := utils.ExtractFileStat(utils.SplitStringAndGetIndexSafely(line, " ", 1))
			var log = pb.Log{
				Type:         "错误",
				Size:         size,
				AccessTime:   accessTime,
				ModifiedTime: modifyTime,
			}
			site.Logs = append(site.Logs, &log)
		}
		if strings.Contains(line, "CustomLog") {
			size, accessTime, modifyTime := utils.ExtractFileStat(utils.SplitStringAndGetIndexSafely(line, " ", 1))
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
		option = filepath.Join(rootPath, option)
		if strings.Contains(option, "*") {

			files, err := utils.FindMatchedFiles(option)
			if err != nil {
				return err
			}
			for _, file := range files {
				err := recursiveInsertData(file, rootPath, response, envMap)
				if err != nil {
					return err
				}
			}

		} else {
			err := recursiveInsertData(option, rootPath, response, envMap)
			if err != nil {
				return err
			}
		}
	}
	return
}
