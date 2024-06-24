//go:build windows

package get_sys_info

import (
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"go-agent/utils/windows"
	"strings"
)

func platformGetSysInfo() (cpu, cpuProcessor string, loadAverage *pb.LoadAverage, err error) {
	cpu, err = windows.RunCmd("wmic cpu get name | findstr /v 'Name'")
	cpu = utils.SplitStringAndGetIndexSafely(cpu, "\n", 1)
	cpu = strings.Trim(cpu, "\r\n ")

	cpuProcessor, err = windows.RunCmd("wmic cpu get NumberOfCores | findstr /v 'NumberOfCores'")
	cpuProcessor = utils.SplitStringAndGetIndexSafely(cpuProcessor, "\n", 1)
	cpuProcessor = strings.Trim(cpuProcessor, "\r\n ")
	//windows没有load average,默认为 nil
	loadAverage = nil

	return
}
