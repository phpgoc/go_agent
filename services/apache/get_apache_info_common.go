package apache

import (
	"context"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"path/filepath"
	"strconv"
	"strings"
)

type Server struct {
	pb.UnimplementedApacheServiceServer
}

func (s *Server) GetApacheInfo(_ context.Context, _ *pb.GetApacheInfoRequest) (*pb.GetApacheInfoResponse, error) {
	utils.LogInfo("called ApacheInfo")
	var response pb.GetApacheInfoResponse
	var err error
	var apacheV string

	apacheCmd := utils.FindCommandFromPathAndProcessByMatchStringArray([]string{"apache2", "httpd", "apache"})
	if apacheCmd == "" {
		response.Message = "can't find apache"
		utils.LogError(response.Message)
		return &response, err
	}

	apacheV, err = utils.RunCmd(apacheCmd + " -V")
	if err != nil {
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
	envContent := utils.GetFirstAndLogError(
		func() (string, error) {
			return utils.ReadFile(filepath.Join(httpRoot, "envvars"))
		})
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
			parseInt := utils.GetFirstAndLogError(
				func() (int64, error) {
					return strconv.ParseInt(listenString, 10, 32)
				})
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
			filePath := utils.SplitStringAndGetIndexSafely(line, " ", 1)
			filePath = utils.ReplaceStrUseEnvMapStrictWithBrace(filePath, envMap)
			size, accessTime, modifyTime := utils.ExtractFileStat(filePath)
			var log = pb.Log{
				Type:         "错误",
				FilePath:     filePath,
				Size:         size,
				AccessTime:   accessTime,
				ModifiedTime: modifyTime,
			}
			site.Logs = append(site.Logs, &log)
		}
		if strings.Contains(line, "CustomLog") {
			filePath := utils.SplitStringAndGetIndexSafely(line, " ", 1)
			filePath = utils.ReplaceStrUseEnvMapStrictWithBrace(filePath, envMap)
			size, accessTime, modifyTime := utils.ExtractFileStat(filePath)
			var log = pb.Log{
				Type:         "访问",
				FilePath:     filePath,
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
