package apache

import (
	"context"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"path/filepath"
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
	var httpDefaultRoot, serverDefaultConfig string
	for _, line := range strings.Split(apacheV, "\n") {
		if strings.Contains(line, "HTTPD_ROOT") {
			if httpDefaultRoot = utils.SplitStringAndGetIndexSafely(line, "=", 1); httpDefaultRoot != "" {
				httpDefaultRoot = strings.Trim(httpDefaultRoot, "\"")

			} else {
				break
			}
		}
		if strings.Contains(line, "SERVER_CONFIG_FILE") {
			if serverDefaultConfig = utils.SplitStringAndGetIndexSafely(line, "=", 1); serverDefaultConfig != "" {
				serverDefaultConfig = strings.Trim(serverDefaultConfig, "\"")
				//set empty string to split
			} else {
				break
			}
		}
	}
	if httpDefaultRoot == "" || serverDefaultConfig == "" {
		//如果能找到命令,不可能找不到这两个'
		response.Message = "can't find HTTPD_ROOT or SERVER_CONFIG_FILE"
		utils.LogError(response.Message)
		return &response, err
	}
	//获取环境变量里的值 得到一个map
	envMap := utils.GetSystemEnvVars()

	KVMap := platformReadEnvFile(httpDefaultRoot, envMap)

	err = insertApacheInstance(filepath.Join(httpDefaultRoot, serverDefaultConfig), &response, KVMap)
	if err != nil {
		utils.LogError(err.Error())
	}
	utils.LogInfo(response.String())
	return &response, nil
}

func insertApacheInstance(configFileName string, response *pb.GetApacheInfoResponse, envMap map[string]string) (err error) {

	return
}
