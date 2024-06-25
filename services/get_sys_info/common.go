package get_sys_info

import (
	"context"
	"github.com/elastic/go-sysinfo"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"regexp"
	"strconv"
	"strings"
)

type GetSysInfoServer struct {
	pb.UnimplementedGetSysInfoServer
}

func (s *GetSysInfoServer) GetSysInfo(_ context.Context, _ *pb.GetSysInfoRequest) (*pb.GetSysInfoResponse, error) {
	host, err := sysinfo.Host()
	var res pb.GetSysInfoResponse
	if err != nil {
		res.Message = err.Error()
		return &res, err
	}
	var info = host.Info()
	memory, err := host.Memory()
	if err != nil {
		res.Message = err.Error()
		return &res, err
	}

	//var loadAverage = host.LoadAverage()
	//Caption              string   `protobuf:"bytes,2,opt,name=caption,proto3" json:"caption,omitempty"`
	//	Timezone             string   `protobuf:"bytes,3,opt,name=timezone,proto3" json:"timezone,omitempty"`
	//	SysVersion           string   `protobuf:"bytes,4,opt,name=sysVersion,proto3" json:"sysVersion,omitempty"`
	//	SysType              string   `protobuf:"bytes,5,opt,name=sysType,proto3" json:"sysType,omitempty"`
	//	SysArch              string   `protobuf:"bytes,6,opt,name=sysArch,proto3" json:"sysArch,omitempty"`
	//	Hostname             string   `protobuf:"bytes,7,opt,name=hostname,proto3" json:"hostname,omitempty"`
	//	Uptime               string   `protobuf:"bytes,8,opt,name=uptime,proto3" json:"uptime,omitempty"`
	//	BootTime             string   `protobuf:"bytes,9,opt,name=bootTime,proto3" json:"bootTime,omitempty"`
	//	LastShutdownTime     string   `protobuf:"bytes,10,opt,name=lastShutdownTime,proto3" json:"lastShutdownTime,omitempty"`
	//	Cpu                  string   `protobuf:"bytes,11,opt,name=cpu,proto3" json:"cpu,omitempty"`
	//	CpuProcessor         string   `protobuf:"bytes,12,opt,name=cpuProcessor,proto3" json:"cpuProcessor,omitempty"`
	//	LoadAverage          []string `protobuf:"bytes,13,rep,name=loadAverage,proto3" json:"loadAverage,omitempty"`
	//	Memory               string   `protobuf:"bytes,14,opt,name=memory,proto3" json:"memory,omitempty"`
	//	Ips

	//platform.platform() windows or linux
	res.Caption = info.OS.Platform
	offset := info.TimezoneOffsetSec / 3600
	var offsetString = strconv.Itoa(offset)
	if offset >= 0 {
		offsetString = "+" + offsetString
	}

	res.Timezone = offsetString
	res.SysVersion = info.OS.Version
	res.SysType = info.OS.Type
	res.SysArch = info.Architecture
	res.Hostname = info.Hostname
	res.BootTime = utils.FormatTime(info.BootTime)
	//res.LastShutdownTime = utils.FormatTime(info.LastShutdownTime)
	res.Uptime = utils.FormatDuration(info.Uptime())
	res.BootTime = utils.FormatTime(info.BootTime)

	res.Cpu, res.CpuProcessor, res.LoadAverage, err = platformGetSysInfo()
	if err != nil {
		res.Message = err.Error()
		return &res, err
	}

	res.Memory = utils.FormatBytes(memory.Total)

	//ip regex pattern
	re, _ := regexp.Compile(`\d+\.\d+\.\d+\.\d+/\d{1,2}`)
	for _, ip := range info.IPs {
		if re.MatchString(ip) && !strings.HasPrefix(ip, "127.0.0.1") {
			res.Ips = append(res.Ips, ip)
		}
	}
	//留着吧
	utils.LogInfo(res.String())

	return &res, nil
}
