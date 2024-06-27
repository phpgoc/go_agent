package get_sys_info

import (
	"context"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"strconv"
	"time"
)

type GetSysInfoServer struct {
	pb.UnimplementedGetSysInfoServer
}

func (s *GetSysInfoServer) GetSysInfo(_ context.Context, _ *pb.GetSysInfoRequest) (*pb.GetSysInfoResponse, error) {
	//hostInfo, err := sysinfo.Host()
	hostInfo, err := host.Info()
	var res pb.GetSysInfoResponse
	if err != nil {
		//这个找不到确实没法继续了
		res.Message = err.Error()
		return &res, err
	}

	platform, family, version, err := host.PlatformInformation()
	res.Platform = &pb.PlatformModel{
		Platform: platform,
		Family:   family,
		Version:  version,
	}
	_, offset := time.Now().Zone()
	println(offset)
	offset = offset / 3600
	if offset >= 0 {
		res.Timezone = "+" + strconv.Itoa(offset)
	} else {
		res.Timezone = strconv.Itoa(offset)
	}
	res.KernelVersion = hostInfo.KernelVersion
	res.Os = hostInfo.OS
	res.Arch = hostInfo.KernelArch
	res.Hostname = hostInfo.Hostname
	res.BootTime = utils.FormatTimeByTimestamp(int64(hostInfo.BootTime))
	uptime, err := host.Uptime()
	if err != nil {
		utils.LogError(err.Error())
	} else {
		println(uptime)
		res.Uptime = utils.FormatDuration(time.Duration(uptime) * time.Second)
	}
	////res.LastShutdownTime = utils.FormatTime(info.LastShutdownTime)
	//res.Uptime = utils.FormatDuration(info.Uptime())
	//res.BootTime = utils.FormatTime(info.BootTime)

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

	//留着吧
	utils.LogInfo(res.String())

	return &res, nil
}
