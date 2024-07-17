package system

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"strconv"
	"time"
)

func (s *Server) GetSysInfo(_ context.Context, _ *pb.GetSysInfoRequest) (*pb.GetSysInfoResponse, error) {

	hostInfo, err := host.Info()
	var res pb.GetSysInfoResponse
	if err != nil {
		//这个找不到确实没法继续了
		res.Message = err.Error()
		utils.LogError(err.Error())
		return &res, err
	}

	platform, family, version, err := host.PlatformInformation()
	if err != nil {
		utils.LogError(err.Error())
	} else {
		res.Platform = &pb.PlatformModel{
			Platform: platform,
			Family:   family,
			Version:  version,
		}
	}

	_, offset := time.Now().Zone()

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
		res.Uptime = utils.FormatDuration(time.Duration(uptime) * time.Second)
	}

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

	//windows 在启动初期会返回0
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
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		utils.LogError(err.Error())
	} else {
		res.VirtualMemory = &pb.MemoryStat{
			Total:     utils.FormatBytes(virtualMemory.Total),
			Available: utils.FormatBytes(virtualMemory.Available),
			Used:      utils.FormatBytes(virtualMemory.Used),
			Free:      utils.FormatBytes(virtualMemory.Free),
		}
	}
	swapMemory, err := mem.SwapMemory()
	if err != nil {
		utils.LogError(err.Error())
	} else {
		res.SwapMemory = &pb.MemoryStat{
			Total:     utils.FormatBytes(swapMemory.Total),
			Used:      utils.FormatBytes(swapMemory.Used),
			Free:      utils.FormatBytes(swapMemory.Free),
			Available: "swap have no available",
		}
	}
	//只需要 物理硬盘
	//如果需要 cd-rom drives, USB keys 等, 使用true
	disks, err := disk.Partitions(false)
	if err != nil {
		utils.LogError(err.Error())
	} else {
		utils.LogInfo(fmt.Sprintf("disks.len: %d ", len(disks)))
		for _, d := range disks {
			diskStat, err := disk.Usage(d.Mountpoint)
			if err != nil {
				utils.LogError(err.Error())
			} else {
				res.Disks = append(res.Disks, &pb.Disk{
					Device:     d.Device,
					MountPoint: d.Mountpoint,
					FsType:     d.Fstype,
					Usage: &pb.DiskUsage{
						Total: utils.FormatBytes(diskStat.Total),
						Used:  utils.FormatBytes(diskStat.Used),
						Free:  utils.FormatBytes(diskStat.Free),
						//除零会显示为NAN%
						UsedPercent: fmt.Sprintf("%.3f%%", float64(diskStat.Used*100)/float64(diskStat.Total)),
					},
				})
			}
		}

	}

	//留着吧
	utils.LogInfo(res.String())

	return &res, nil
}
