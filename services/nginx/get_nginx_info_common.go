package nginx

import (
	"context"
	pb "go-agent/agent_proto"
	"go-agent/runtime"
	"go-agent/utils"
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
	utils.LogWarn(commandPath)
	nginxV, err := utils.RunCmd(commandPath + " -V")
	utils.LogWarn(nginxV)
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

	re, _ = regexp.Compile(`-c\s+(\S+)`)

	for _, process := range runtime.Processes {

		exe, _ := process.Exe()

		if exe == commandPath {
			cmd, _ := process.Cmdline()
			//能找到执行这个命令时是在哪个目录下吗

			if matched := re.FindStringSubmatch(cmd); len(matched) > 1 {
				if defaultConfigFile == matched[1] {
					continue
				}
				if utils.IsAbsolutePath(matched[1]) {
					insertNginxInfo(&commandPath, matched[1], &res)
				} else {
					runCmdDIr, _ := process.Cwd()
					insertNginxInfo(&commandPath, runCmdDIr+matched[1], &res)
				}
			}
		}

	}
	_ = defaultErrorLog
	_ = defaultAccessLog
	return &res, nil
}

func insertNginxInfo(commandPath *string, configFile string, res *pb.GetNginxInfoResponse) {
	// do something
}
