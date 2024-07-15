package agent_runtime

import (
	"os"
	"path/filepath"
)

var (
	defaultLogFileNameAtTempDir = "xxx/agent.log"
	workDir, _                  = os.Getwd()
	OutDir                      = filepath.Join(workDir, "out")
	//OutDir = "/tmp/out" //测试用
)
