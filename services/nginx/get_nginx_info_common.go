package nginx

import (
	"context"
	"fmt"
	"github.com/tufanbarisyildirim/gonginx"
	"github.com/tufanbarisyildirim/gonginx/parser"
	pb "go-agent/agent_proto"
	"go-agent/agent_runtime"
	utils "go-agent/utils"
	"path/filepath"
	"regexp"
	"strings"
)

var defaultConfigFile, defaultErrorLog, defaultAccessLog string

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
	//库还不成熟,include的文件自动解析会有问题
	//p, err := parser.NewParser(configFile, parser.WithIncludeParsing())
	p, err := parser.NewParser(configFile)
	if err != nil {
		utils.LogError(err.Error())
		return
	}
	parsed, err := p.Parse()
	if err != nil {
		utils.LogError(err.Error())
		return
	}
	var thisNginxInstance pb.NginxInstance
	var includeDirectives []gonginx.IDirective
	var httpErrorlog = defaultErrorLog
	var httpAccesslog = defaultAccessLog
	for _, pi := range parsed.Directives {
		if pi.GetName() == "http" {
			for _, pi2 := range pi.GetBlock().GetDirectives() {
				switch pi2.GetName() {
				case "error_log":
					//string数组
					httpErrorlog = pi2.GetParameters()[0]
				case "access_log":
					//string数组
					httpAccesslog = pi2.GetParameters()[0]
				case "include":
					if include, ok := pi2.(*gonginx.Include); ok {
						utils.LogInfo(include.IncludePath)
						if strings.Contains(include.IncludePath, "site") {

							files, _ := utils.FindMatchedFiles(include.IncludePath)
							for _, file := range files {
								p3, err := parser.NewParser(file)
								if err != nil {
									utils.LogError(err.Error())
									return
								}
								parsed3, err := p3.Parse()
								if err != nil {
									utils.LogError(err.Error())
									return
								}
								includeDirectives = append(includeDirectives, parsed3.Directives...)
							}
						} else if strings.HasSuffix(include.IncludePath, ".conf") {
							//把解析的东西放到parsed里不会有问题吧?
							files, _ := utils.FindMatchedFiles(include.IncludePath)
							for _, file := range files {
								p3, err := parser.NewParser(file)
								if err != nil {
									utils.LogError(err.Error())
									return
								}
								parsed3, err := p3.Parse()
								if err != nil {
									utils.LogError(err.Error())
									return
								}
								includeDirectives = append(includeDirectives, parsed3.Directives...)

							}
						} else {
							continue
						}
					}
				case "server":

				default:
					//utils.LogInfo(pi2.GetName())
				}
			}
		}
	}

	for _, pi := range includeDirectives {
		switch pi.GetName() {
		case "server":
			println("server")
		default:
			utils.LogInfo(pi.GetName())
		}
	}
	size, _, modifyTime := utils.ExtractFileStat(httpErrorlog)
	thisNginxInstance.ErrorLog = &pb.NginxLog{
		FilePath:   httpErrorlog,
		Size:       utils.FormatBytes(size),
		ModifyTime: modifyTime,
	}
	size, _, modifyTime = utils.ExtractFileStat(httpAccesslog)
	thisNginxInstance.AccessLog = &pb.NginxLog{
		FilePath:   httpAccesslog,
		Size:       utils.FormatBytes(size),
		ModifyTime: modifyTime,
	}

	res.NginxInstances = append(res.NginxInstances, &thisNginxInstance)
	// do something
}
