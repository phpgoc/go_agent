package agent_runtime

import "github.com/shirou/gopsutil/v4/process"

// 默认每次都获取,除非设置true
var processes, _ = process.Processes()
