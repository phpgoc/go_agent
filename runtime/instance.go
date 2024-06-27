package runtime

import "github.com/shirou/gopsutil/v4/process"

// 只在agent进程初始化时录入，如果需要时时更新的进程需要重新获取，一般不需要
var Processes []*process.Process
