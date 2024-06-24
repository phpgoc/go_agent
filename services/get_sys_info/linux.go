//go:build linux

package get_sys_info

import (
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"go-agent/utils/linux"
	"strings"
)

func platformGetSysInfo() (cpu, cpuProcessor string, loadAverage *pb.LoadAverage, err error) {
	cpu, _ = linux.RunCmd("cat /proc/cpuinfo | grep 'model name' | uniq | awk -F: '{print $2}'")
	cpuProcessor, _ = linux.RunCmd("cat /proc/cpuinfo | grep 'processor' | wc -l")

	cpu = strings.Trim(cpu, "\n")
	cpuProcessor = strings.Trim(cpuProcessor, "\n")

	loadAverageString, _ := linux.RunCmd("cat /proc/loadavg")
	loadAverage = &pb.LoadAverage{
		One:     utils.SplitStringAndGetIndexSafely(loadAverageString, " ", 0),
		Five:    utils.SplitStringAndGetIndexSafely(loadAverageString, " ", 1),
		Fifteen: utils.SplitStringAndGetIndexSafely(loadAverageString, " ", 2),
	}
	return
}
