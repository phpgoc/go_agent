package apache

import (
	"context"
	"errors"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"path/filepath"
	"regexp"
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
		//如果能找到命令,不可能找不到这两个
		response.Message = "can't find HTTPD_ROOT or SERVER_CONFIG_FILE"
		utils.LogError(response.Message)
		return &response, err
	}
	//获取环境变量里的值 得到一个map
	envMap := utils.GetSystemEnvVars()

	KVMap := platformReadEnvFile(httpDefaultRoot, envMap)

	err = insertApacheInstance(filepath.Join(httpDefaultRoot, serverDefaultConfig), httpDefaultRoot, &response, KVMap)
	if err != nil {
		utils.LogError(err.Error())
	}
	utils.LogInfo(response.String())
	return &response, nil
}

// insertApacheInstance apache的配置文件加载过于复杂,目前不支持多个apache实例
func insertApacheInstance(configFileName, httpdRoot string, response *pb.GetApacheInfoResponse, envMap map[string]string) (err error) {
	inVirtualHost := false
	var thisVirtualHost = pb.ApacheVirtualHost{}
	//目前这行看起来没用,只为用这个变量复制,未来如果有可能支持多个apache实例时需要改的代码少写
	thisInstance := response

	configFileContent, err := utils.ReadFile(configFileName)

	sourceConfig2ContentMap := map[string]string{configFileName: configFileContent}
	includeConfig2ContentMap := make(map[string]string)
	for {
		if len(sourceConfig2ContentMap) == 0 {
			break
		}
		for file, content := range sourceConfig2ContentMap {
			var useful = false
			lines := strings.Split(content, utils.LineBreak)
			for lineI, line := range lines {

				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "#") {
					continue
				}
				if strings.Contains(line, `$`) {
					line = utils.ReplaceStrUseEnvMapStrictWithBrace(line, envMap)
				}
				if strings.HasPrefix(line, "<VirtualHost") {
					inVirtualHost = true
					thisVirtualHost = pb.ApacheVirtualHost{}
					re := regexp.MustCompile(`<VirtualHost(?:\s+(\S+))+>`)
					matched := re.FindStringSubmatch(line)
					if len(matched) < 2 {
						utils.LogError("VirtualHost format error")
						return
					}
					for i := 1; i < len(matched); i++ {
						thisVirtualHost.Listens = append(thisVirtualHost.Listens, matched[i])
					}
				} else if strings.HasPrefix(line, "</VirtualHost") {
					if !inVirtualHost {
						utils.LogError("/VirtualHost format error")
						return
					}
					thisInstance.VirtualHosts = append(thisInstance.VirtualHosts, &thisVirtualHost)
					inVirtualHost = false
					thisVirtualHost = pb.ApacheVirtualHost{}
				} else if strings.HasPrefix(line, "Include") {
					re := regexp.MustCompile(`Include(?:Optional)?(?:\s+(\S+))+`)
					matched := re.FindStringSubmatch(line)
					if len(matched) < 2 {
						utils.LogError("Include format error")
						//不致命
						continue
					}
					for i := 1; i < len(matched); i++ {
						matchedI := matched[i]
						if !utils.IsAbsolutePath(matchedI) {
							matchedI = filepath.Join(httpdRoot, matchedI)
						}
						if strings.Contains(matched[i], "*") {

							includeFiles, _ := utils.FindMatchedFiles(matched[i])
							for _, includeFile := range includeFiles {
								//文件一定存在
								includeConfig2ContentMap[includeFile], _ = utils.ReadFile(includeFile)
							}
						} else {
							if !utils.FileExists(matchedI) {
								continue
							}
							includeConfig2ContentMap[matchedI], _ = utils.ReadFile(matchedI)
						}
					}

				} else if strings.HasPrefix(line, "ServerName") {
					re := regexp.MustCompile(`ServerName(?:\s+(\S+))+`)
					matched := re.FindStringSubmatch(line)
					if len(matched) < 2 {
						utils.LogError("ServerName format error")
						return
					}
					for i := 1; i < len(matched); i++ {
						thisVirtualHost.ServerNames = append(thisVirtualHost.ServerNames, matched[i])
					}

				} else if strings.HasPrefix(line, "Listen") {
					if inVirtualHost {
						utils.LogError("Listen cannot occur within <VirtualHost> section")
						return
					}
					listenMatched, err := extractMatchesFromLine(line, "Listen")
					if err != nil {
						utils.LogError(err.Error())
						return err
					}
					thisInstance.Listens = append(thisInstance.Listens, listenMatched...)
				} else if strings.HasPrefix(line, "DocumentRoot") {
					root, err := extractMatchesFromLine(line, "DocumentRoot")
					if err != nil {
						utils.LogError(err.Error())
						return err
					}
					//不应该有多个
					if len(root) != 1 {
						utils.LogError("DocumentRoot format error")
						return err
					}
					// 如果在VirtualHost中,则设置VirtualHost的Root,否则设置Instance的Root
					if inVirtualHost {
						thisVirtualHost.Root = root[0]
					} else {
						thisInstance.Root = root[0]
					}
				} else if strings.HasPrefix(line, "ErrorLog") {
					errLog := ""
					if strings.Contains(line, `"`) {
						//自己读,读到两个引号中间的内容
						re := regexp.MustCompile(`"(.*?)"`)
						matched := re.FindStringSubmatch(line)
						if len(matched) < 2 {
							utils.LogError("ErrorLog format error")
							return
						}

						errLog = matched[1]
					} else {
						matchedErrorLog, err := extractMatchesFromLine(line, "ErrorLog")
						if err != nil {
							utils.LogError(err.Error())
							return err
						}
						errLog = matchedErrorLog[0]
					}

					if !utils.IsAbsolutePath(errLog) {
						errLog = filepath.Join(httpdRoot, errLog)
					}

					size, _, modifyTime := utils.ExtractFileStat(errLog)
					log := &pb.ApacheLog{
						Size:       utils.FormatBytes(size),
						ModifyTime: modifyTime,
						FilePath:   errLog,
					}
					if inVirtualHost {
						thisVirtualHost.ErrorLog = log
					} else {
						thisInstance.ErrorLog = log
					}

				} else if strings.HasPrefix(line, "CustomLog") {
					customLog := ""
					if strings.Contains(line, `"`) {
						//自己读,读到两个引号中间的内容
						re := regexp.MustCompile(`"(.*?)"`)
						matched := re.FindStringSubmatch(line)
						if len(matched) < 2 {
							utils.LogError("ErrorLog format error")
							return
						}

						customLog = matched[1]
					} else {
						matchedErrorLog, err := extractMatchesFromLine(line, "ErrorLog")
						if err != nil {
							utils.LogError(err.Error())
							return err
						}
						customLog = matchedErrorLog[0]
					}

					if !utils.IsAbsolutePath(customLog) {
						customLog = filepath.Join(httpdRoot, customLog)
					}
					size, _, modifyTime := utils.ExtractFileStat(customLog)
					log := &pb.ApacheLog{
						Size:       utils.FormatBytes(size),
						ModifyTime: modifyTime,
						FilePath:   customLog,
					}
					if inVirtualHost {
						thisVirtualHost.CustomLog = log
					} else {
						thisInstance.CustomLog = log
					}

				} else {
					continue
				}
				//进入了任何一个if,都是有用的
				useful = true
				_ = lineI
			} // end for _, line := range lines

			if useful {
				thisInstance.ConfigFiles = append(thisInstance.ConfigFiles, file)
			}
		} // end for file, content := range sourceConfig2ContentMap

		sourceConfig2ContentMap = includeConfig2ContentMap
		includeConfig2ContentMap = make(map[string]string)
	}

	return
}

// <pattern> <matched1> [<matched2> ...]
func extractMatchesFromLine(line, pattern string) ([]string, error) {
	re, err := regexp.Compile(pattern + `(?:\s+(\S+))+`)
	if err != nil {
		return nil, err
	}
	matched := re.FindStringSubmatch(line)
	if len(matched) < 2 {
		return nil, errors.New(pattern + " format error")
	}
	return matched[1:], nil
}
