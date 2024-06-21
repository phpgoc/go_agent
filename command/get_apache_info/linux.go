//go:build linux

package get_apache_info

import (
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"go-agent/utils/linux"
	"path/filepath"
	"strings"
)

// 在用windows版本后再开启条件编译
func PlatformGetApacheInfo() (*pb.GetApacheInfoResponse, error) {
	var response pb.GetApacheInfoResponse
	var err error
	var apacheV string

	//大V可能会返回错误，即使有apache，报错还不不影响apache正常运行，所用用小v判断是否存在，如果存在则用大V
	if _, err = linux.RunCmd("apache2 -v"); err == nil {
		apacheV, err = linux.RunCmd("apache2 -V")
	} else if apacheV, err = linux.RunCmd("apache -v"); err == nil {
		apacheV, _ = linux.RunCmd("apache -V")
	} else if apacheV, err = linux.RunCmd("httpd -v"); err == nil {
		apacheV, _ = linux.RunCmd("httpd -V")
	} else {
		response.Message = "can't find apache"
		return &response, err
	}
	//println(apacheV)
	//find HTTPD_ROOT and SERVER_CONFIG_FILE from apache -V
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
	err = recursiveInsertData(file, httpRoot, response, envMap)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
