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
	commandPath, _ := utils.FindCommandFromPathAndProcessByMatchStringArray([]string{"nginx"})
	if commandPath == "" {
		utils.LogError("can't find nginx")
		return nil, nil
	}
	var res pb.GetNginxInfoResponse
	var prefix string
	nginxV, err := utils.RunCmd(commandPath + " -V")
	if err != nil {
		utils.LogError(err.Error())
		return nil, nil
	}

	var reC *regexp.Regexp
	reC, _ = regexp.Compile(`--conf-path=(\S+)`)
	if matched := reC.FindStringSubmatch(nginxV); len(matched) > 1 {
		defaultConfigFile = matched[1]
	} else {
		//应该没可能到这
		utils.LogError("can't find default config file")
		return nil, nil
	}
	reC, _ = regexp.Compile(`--error-log-path=(\S+)`)
	if matched := reC.FindStringSubmatch(nginxV); len(matched) > 1 {
		defaultErrorLog = matched[1]
	} else {
		utils.LogError("can't find default error log")
		return nil, nil
	}

	reC, _ = regexp.Compile(`--http-log-path=(\S+)`)
	if matched := reC.FindStringSubmatch(nginxV); len(matched) > 1 {
		defaultAccessLog = matched[1]
	} else {
		utils.LogError("can't find default access log")
		return nil, nil
	}
	reC, _ = regexp.Compile(`--prefix=(\S+)`)
	if matched := reC.FindStringSubmatch(nginxV); len(matched) > 1 {
		prefix = matched[1]
	} else {
		utils.LogError("can't find default prefix")
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
	insertNginxInfo(defaultConfigFile, prefix, &res)

	//不指定config的情况上边已经处理了
	reC, _ = regexp.Compile(`-c\s+(\S+)`)

	for _, process := range agent_runtime.GetProcesses() {

		exe, _ := process.Exe()

		if exe == commandPath {
			cmd := utils.GetFirstAndLogError(
				func() (string, error) {
					return process.Cmdline()
				})

			processRunCmdDIr := utils.GetFirstAndLogError(
				func() (string, error) {
					return process.Cwd()
				})

			//进程里的-c必须匹配到,匹配不到的是默认配置,已经处理过了
			if matched := reC.FindStringSubmatch(cmd); len(matched) > 1 {

				thisConfigFile := matched[1]
				if !utils.IsAbsolutePath(thisConfigFile) {
					thisConfigFile = filepath.Join(processRunCmdDIr, matched[1])
				}
				rePrefix := regexp.MustCompile(`-p\s+(\S+)`)
				//prefix用来拼接config里的相对路径,相对路径绝不是互相相对,必须是相对nginx的prefix
				if matchedPrefix := rePrefix.FindStringSubmatch(cmd); len(matchedPrefix) > 1 {

					if utils.IsAbsolutePath(matchedPrefix[1]) {
						insertNginxInfo(thisConfigFile, matchedPrefix[1], &res)
					} else {
						insertNginxInfo(thisConfigFile, filepath.Join(processRunCmdDIr, matchedPrefix[1]), &res)
					}
				} else {
					//没匹配到-p说明用编译时的prefix
					insertNginxInfo(thisConfigFile, prefix, &res)
				}

			}
		}

	}
	utils.LogInfo(fmt.Sprintf("NginxInfo res:%s", InfoResponseWrapper{&res}))
	return &res, nil
}

func insertNginxInfo(configFile string, processPrefix string, res *pb.GetNginxInfoResponse) {
	utils.LogInfo(fmt.Sprintf("configFile:%v", configFile))
	if !utils.IsAbsolutePath(configFile) {
		//能进来这里说明代码写错了,上边没有进行绝对路径的拼接
		//或者测试用例用错了
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
	var thisNginxInstance = pb.NginxInstance{
		ConfigPath: configFile,
	}
	var instanceErrorLog = defaultErrorLog
	var instanceAccessLog = defaultAccessLog
	var includeDirectives []gonginx.IDirective
	var searchDirectives []gonginx.IDirective
	for _, pi := range parsed.Directives {
		if pi.GetName() == "http" {
			searchDirectives = pi.GetBlock().GetDirectives()
			break
		}
	}
	for {
		if searchDirectives == nil {
			break
		}
		for _, pi := range searchDirectives {
			switch pi.GetName() {
			case "error_log":
				if !utils.IsAbsolutePath(pi.GetParameters()[0]) {
					instanceErrorLog = filepath.Join(processPrefix, pi.GetParameters()[0])
				} else {
					instanceErrorLog = pi.GetParameters()[0]
				}
			case "access_log":
				if !utils.IsAbsolutePath(pi.GetParameters()[0]) {
					instanceAccessLog = filepath.Join(processPrefix, pi.GetParameters()[0])
				} else {
					instanceAccessLog = pi.GetParameters()[0]
				}
			case "include":
				if include, ok := pi.(*gonginx.Include); ok {
					includePath := include.IncludePath
					if !utils.IsAbsolutePath(include.IncludePath) {
						includePath = filepath.Join(processPrefix, include.IncludePath)
					}
					utils.LogInfo(includePath)
					if strings.Contains(includePath, "mime") {
						//这个文件不是nginx的配置文件,不需要解析
					} else if strings.HasSuffix(includePath, ".conf") {
						//把解析的东西放到parsed里不会有问题吧?
						files := utils.GetFirstAndLogError(
							func() ([]string, error) {
								return utils.FindMatchedFiles(includePath)
							})
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
						//这里还有很多出错可能,最上边的不寻找mime的条件可能不足以保证正确性
						files := utils.GetFirstAndLogError(
							func() ([]string, error) {
								return utils.FindMatchedFiles(includePath)
							})
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
					}
				}
			case "server":

				var thisServer = pb.NginxServerInfo{}
				for _, serverI := range pi.GetBlock().GetDirectives() {
					switch serverI.GetName() {
					case "server_name":
						thisServer.ServerName = serverI.GetParameters()[0]
					case "listen":

						thisServer.Listens = append(thisServer.Listens, serverI.GetParameters()[0])
					case "root":
						thisServer.Root = serverI.GetParameters()[0]
					case "error_log":
						errorLogFile := serverI.GetParameters()[0]
						if !utils.IsAbsolutePath(serverI.GetParameters()[0]) {
							errorLogFile = filepath.Join(processPrefix, serverI.GetParameters()[0])
						}
						size, _, modifyTime := utils.ExtractFileStat(errorLogFile)
						thisServer.ErrorLog = &pb.NginxLog{
							FilePath:   serverI.GetParameters()[0],
							Size:       utils.FormatBytes(size),
							ModifyTime: modifyTime,
						}
					case "access_log":
						accessLogFile := serverI.GetParameters()[0]
						if !utils.IsAbsolutePath(serverI.GetParameters()[0]) {
							accessLogFile = filepath.Join(processPrefix, serverI.GetParameters()[0])
						}
						size, _, modifyTime := utils.ExtractFileStat(accessLogFile)
						thisServer.AccessLog = &pb.NginxLog{
							FilePath:   serverI.GetParameters()[0],
							Size:       utils.FormatBytes(size),
							ModifyTime: modifyTime,
						}
					default:
						//utils.LogInfo(serverI.GetName())
						//utils.LogInfo(serverI.GetParameters()[0])
						continue
					}
				}
				thisNginxInstance.Servers = append(thisNginxInstance.Servers, &thisServer)
			default:
				//utils.LogInfo(pi2.GetName())
			}

		}
		searchDirectives = includeDirectives
		includeDirectives = nil
	}

	size, _, modifyTime := utils.ExtractFileStat(instanceErrorLog)
	thisNginxInstance.ErrorLog = &pb.NginxLog{
		FilePath:   instanceErrorLog,
		Size:       utils.FormatBytes(size),
		ModifyTime: modifyTime,
	}
	size, _, modifyTime = utils.ExtractFileStat(instanceAccessLog)
	thisNginxInstance.AccessLog = &pb.NginxLog{
		FilePath:   instanceAccessLog,
		Size:       utils.FormatBytes(size),
		ModifyTime: modifyTime,
	}

	res.NginxInstances = append(res.NginxInstances, &thisNginxInstance)
}
