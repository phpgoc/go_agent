package utils

import (
	"go-agent/agent_runtime"
	"os"
	"path/filepath"
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
	if *agent_runtime.UseFileToLog {
		fullFileName := filepath.Join(os.TempDir(), *agent_runtime.LogFileNameAtTempDir)
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
