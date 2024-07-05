package system

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/process"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"golang.org/x/net/context"
	"strconv"
	"strings"
)

func (s *Server) GetProcessList(_ context.Context, req *pb.GetProcessListRequest) (*pb.GetProcessListResponse, error) {
	processList, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var res pb.GetProcessListResponse
	for _, p := range processList {
		processMemorySource := utils.GetFirstAndLogError(func() (*process.MemoryInfoStat, error) {
			return p.MemoryInfo()
		})
		var thisProcessInfo = &pb.ProcessInfo{
			Pid: p.Pid,
			Ppid: utils.GetFirstAndLogError(func() (int32, error) {
				return p.Ppid()
			}),
			Name: utils.GetFirstAndLogError(func() (string, error) {
				return p.Name()
			}),
			Status: strings.Join(utils.GetFirstWithoutLogError(func() ([]string, error) {
				return p.Status()
			}), ","),
			Username: utils.GetFirstAndLogError(func() (string, error) {
				return p.Username()
			}),
			CreateTime: utils.FormatTimeByTimestamp(
				utils.GetFirstAndLogError(func() (int64, error) {
					return p.CreateTime()
				}, 0)),
			CpuPercent: strconv.FormatFloat(utils.GetFirstAndLogError(func() (float64, error) {
				return p.CPUPercent()
			}),
				'f', -1, 64),
			MemoryPercent: strconv.FormatFloat(float64(utils.GetFirstAndLogError(func() (float32, error) {
				return p.MemoryPercent()
			})),
				'f', -1, 64),

			Exe: utils.GetFirstAndLogError(func() (string, error) {
				return p.Exe()
			}),
			Cwd: utils.GetFirstAndLogError(func() (string, error) {
				return p.Cwd()
			}),
			Cmdline: utils.GetFirstAndLogError(func() (string, error) {
				return p.Cmdline()
			}),
		}
		if processMemorySource != nil {
			thisProcessInfo.Memory = &pb.ProcessMemoryInfoStat{
				VMS:   utils.FormatBytes(processMemorySource.VMS),
				HWM:   utils.FormatBytes(processMemorySource.HWM),
				Data:  utils.FormatBytes(processMemorySource.Data),
				Stack: utils.FormatBytes(processMemorySource.Stack),
			}
		}
		if req.WithThreadTimes {
			threadTimes, _ := p.Threads()
			for tid, threadTimeInfo := range threadTimes {
				thisProcessInfo.Threads = append(thisProcessInfo.Threads, &pb.ThreadTimesStat{
					Tid:     tid,
					User:    threadTimeInfo.User,
					System:  threadTimeInfo.System,
					Idle:    threadTimeInfo.Idle,
					Nice:    threadTimeInfo.Nice,
					IoWait:  threadTimeInfo.Iowait,
					Irq:     threadTimeInfo.Irq,
					SoftIrq: threadTimeInfo.Softirq,
				})
			}
		}
		res.List = append(res.List, thisProcessInfo)
	}
	utils.LogInfo(fmt.Sprintf("processList.len: %d", len(processList)))
	return &res, nil
}
