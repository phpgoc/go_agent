package nginx

import (
	"context"
	"fmt"
	pb "go-agent/agent_proto"
	"go-agent/agent_runtime"
	"go-agent/utils"
	"path/filepath"
	"regexp"
)

func (*Server) GetNginxInfo(_ context.Context, _ *pb.GetNginxInfoRequest) (*pb.GetNginxInfoResponse, error) {
	utils.LogInfo("called NginxInfo")
	commandPath := utils.FindCommandFromPathAndProcessByMatchStringArray([]string{"nginx"})
	if commandPath == "" {
		utils.LogError("can't find nginx")
		return nil, nil
	}
	var res pb.GetNginxInfoResponse
	nginxV, err := utils.RunCmd(commandPath + " -V")
	if err != nil {
		utils.LogError(err.Error())
		return nil, nil
	}
	var defaultConfigFile, defaultErrorLog, defaultAccessLog string

	var re *regexp.Regexp
	re, _ = regexp.Compile(`--conf-path=(\S+)`)
	if matched := re.FindStringSubmatch(nginxV); len(matched) > 1 {
		defaultConfigFile = matched[1]
	} else {
		//应该没可能到这
		utils.LogError("can't find default config file")
		return nil, nil
	}
	re, _ = regexp.Compile(`--error-log-path=(\S+)`)
	if matched := re.FindStringSubmatch(nginxV); len(matched) > 1 {
		defaultErrorLog = matched[1]
	} else {
		utils.LogError("can't find default error log")
		return nil, nil
	}

	re, _ = regexp.Compile(`--http-log-path=(\S+)`)
	if matched := re.FindStringSubmatch(nginxV); len(matched) > 1 {
		defaultAccessLog = matched[1]
	} else {
		utils.LogError("can't find default access log")
		return nil, nil
	}
	//windows的默认config大概是相对路径
	if !utils.IsAbsolutePath(defaultConfigFile) {
		//上边三个都会是相对路径
		commandPathDir := filepath.Dir(commandPath)
		defaultConfigFile = filepath.Join(commandPathDir, defaultConfigFile)
		defaultErrorLog = filepath.Join(commandPathDir, defaultErrorLog)
		defaultAccessLog = filepath.Join(commandPathDir, defaultAccessLog)
	}
	insertNginxInfo(defaultConfigFile, &res)

	//不指定config的情况上边已经处理了
	re, _ = regexp.Compile(`-c\s+(\S+)`)

	for _, process := range agent_runtime.GetProcesses() {

		exe, _ := process.Exe()

		if exe == commandPath {
			cmd, _ := process.Cmdline()
			//能找到执行这个命令时是在哪个目录下吗

			if matched := re.FindStringSubmatch(cmd); len(matched) > 1 {
				if defaultConfigFile == matched[1] {
					continue
				}
				if utils.IsAbsolutePath(matched[1]) {
					insertNginxInfo(matched[1], &res)
				} else {
					runCmdDIr, _ := process.Cwd()
					insertNginxInfo(runCmdDIr+matched[1], &res)
				}
			}
		}

	}
	_ = defaultErrorLog
	_ = defaultAccessLog
	return &res, nil
}

func insertNginxInfo(configFile string, res *pb.GetNginxInfoResponse) {
	utils.LogInfo(fmt.Sprintf("configFile:%v", configFile))
	if !utils.IsAbsolutePath(configFile) {
		//能进来这里只有两个可能
		//代码写错了,上边没有进行绝对路径的拼接
		//nginx 的config是相对路径,这种我完全不知道怎么处理,没有nginx root的说法
		utils.LogError("configFile is not absolute path : " + configFile)
		return
	}
	// do something
}
