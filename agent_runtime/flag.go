package agent_runtime

import (
	"flag"
	"fmt"
)

var UseFileToLog = flag.Bool("use_file", false, fmt.Sprintf("Use file to log, if not set is Std Out,if set default is %s", defaultLogFileNameAtTempDir))

var LogFileNameAtTempDir = flag.String("log_file", defaultLogFileNameAtTempDir, "Log file name,, Must use use_file to take effect,it is about relative paths to temporary folders")

var Port = flag.Int("port", 50051, "The server port")

// 设置为true不再时时获取进程,每次都是用该进程初始化时的状态,性能会有所提升,
var getProcessOnce = flag.Bool("get_process_once", false, "set true to get process every time")
