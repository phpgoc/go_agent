//go:build linux

package get_apache_info

import (
	pb "go-agent/agent_proto"
	"go-agent/utils/linux"
	"path/filepath"
	"strings"
)

// 在用windows版本后再开启条件编译
func PlatformGetApacheInfo() (*pb.GetApacheInfoResponse, error) {
	var response pb.GetApacheInfoResponse
	var err error
	var apacheV string

	//大V可能会报错，即使有apache，报错还不不影响apache正常运行，所用用小v判断是否存在，如果存在则用大V
	if _, err = linux.RunCmd("apache2 -v"); err == nil {
		apacheV, err = linux.RunCmd("apache2 -V")
	} else {
		if apacheV, err = linux.RunCmd("apache -v"); err != nil {
			return nil, err
		} else {
			apacheV, _ = linux.RunCmd("apache -V")
		}
	}
	//println(apacheV)
	//find HTTPD_ROOT and SERVER_CONFIG_FILE from apache -V
	var httpRoot, serverConfig string
	var splitSting []string

	for _, line := range strings.Split(apacheV, "\n") {
		if strings.Contains(line, "HTTPD_ROOT") {
			if splitSting = strings.Split(line, "="); len(splitSting) > 1 {
				httpRoot = strings.Trim(splitSting[1], "\"")
				splitSting = []string{}
				{
				}
			} else {
				break
			}
		}
		if strings.Contains(line, "SERVER_CONFIG_FILE") {
			if splitSting = strings.Split(line, "="); len(splitSting) > 1 {
				serverConfig = strings.Trim(splitSting[1], "\"")
				//set empty string to split
				splitSting = []string{}
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
	println(file)
	recursiveInsertData(file, response)

	return &response, nil
}
