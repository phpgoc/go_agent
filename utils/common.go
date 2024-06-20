package utils

import (
	"fmt"
	"os"
	"path"
	"time"
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

func ExtractFileStat(file string) (size uint64, accessTime, modifyTime string) {
	fi, err := os.Stat(file)
	if err != nil {
		return
	}
	size = uint64(fi.Size())
	accessTime = fi.ModTime().String()
	modifyTime = fi.ModTime().String()
	return
}

func FormatTime(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}
