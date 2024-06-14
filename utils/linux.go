// +build:linux
package utils

import (
	"go-agent/utils/linux"
	"os"
	"path"
)

func Init() (err error) {
	//use LogFileName get path
	dirName := path.Dir(linux.LogFileName)
	//make dir
	err = os.MkdirAll(dirName, 0755)
	if err != nil {
		return err
	}

	linux.LogFile, err = os.OpenFile(linux.LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	return err
}

func LogInfo(log string) (err error) {
	return linux.WriteLogFile(log, "INFO")
}

func LogWarn(log string) (err error) {
	return linux.WriteLogFile(log, "WARN")
}

func LogError(log string) (err error) {
	return linux.WriteLogFile(log, "ERROR")
}
