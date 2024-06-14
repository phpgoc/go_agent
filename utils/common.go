package utils

import (
	"fmt"
	"os"
	"path"
)

var logFile *os.File

func Init() (err error) {
	//use LogFileName get path
	dirName := path.Dir(logFileName)
	//make dir
	err = os.MkdirAll(dirName, 0755)
	if err != nil {
		return err
	}

	logFile, err = os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	osInit()

	return err
}

func writeLogFile(log string, level string) (err error) {
	// log to file
	_, err = logFile.WriteString(fmt.Sprintf("%s %s\n", level, log))
	return err
}

func LogInfo(log string) (err error) {
	return writeLogFile(log, "INFO")
}

func LogWarn(log string) (err error) {
	return writeLogFile(log, "WARN")
}

func LogError(log string) (err error) {
	return writeLogFile(log, "ERROR")
}
