package get_sys_info

import (
	"context"
	"github.com/elastic/go-sysinfo"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/load"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"strconv"
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

	cpuInfo, err := cpu.Info()
	if err != nil {
		utils.LogError(err.Error())
	} else if len(cpuInfo) == 0 {
		utils.LogError("No CPU info found")
	} else {
		logicCount, err := cpu.Counts(true)
		if err != nil {
			utils.LogError(err.Error())
		}
		physicalCount, err := cpu.Counts(false)
		if err != nil {
			utils.LogError(err.Error())
		}
		res.CpuModel = &pb.CpuModel{
			ModelName:     cpuInfo[0].ModelName,
			PhysicalCount: int32(physicalCount),
			LogicalCount:  int32(logicCount),
		}
	}

	loadAverage, err := load.Avg()
	if err != nil {
		utils.LogError(err.Error())
	} else {
		res.LoadAverage = &pb.LoadAverage{
			One:     loadAverage.Load1,
			Five:    loadAverage.Load5,
			Fifteen: loadAverage.Load15,
		}
	}

	res.Memory = utils.FormatBytes(memory.Total)

	//留着吧
	utils.LogInfo(res.String())

	return &res, nil
}
