package utils

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	defaultLogFileNameAtTempDir = "xxx/agent.log"
	useFileToLog                = flag.Bool("use_file", false, fmt.Sprintf("Use file to log, if not set is Std Out,if set default is %s", defaultLogFileNameAtTempDir))
	logFileNameAtTempDir        = flag.String("log_file", defaultLogFileNameAtTempDir, "Log file name,, Must use use_file to take effect,it is about relative paths to temporary folders")
)
var logFile *os.File

// default std out
var writer = os.Stdout

func Init() (err error) {
	//必须是管理员权限，开发初期先注释
	//err = osInitBefore()
	//if err != nil {
	//	return err
	//}
	defer osInitAfter()
	//use LogFileName get path
	if *useFileToLog {
		fullFileName := filepath.Join(os.TempDir(), *logFileNameAtTempDir)
		dirName := filepath.Dir(fullFileName)
		err = os.MkdirAll(dirName, 0755)
		if err != nil {
			return err
		}
		logFile, err = os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		// 适当的条件下设置writer = logFile,默认是os.Stdout
		writer = logFile
	}

	return err
}
