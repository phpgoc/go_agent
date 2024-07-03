package agent_runtime

import (
	"github.com/shirou/gopsutil/v4/process"
	"log"
)

func GetProcesses() []*process.Process {
	if !*getProcessOnce {
		//这里不能使用utils否者会导致循环引用
		log.Print("get process every time")
		processes, _ = process.Processes()
	}
	return processes
}
