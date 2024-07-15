package database

import (
	"bytes"
	"errors"
	"fmt"
	"go-agent/agent_runtime"
	"go-agent/utils"
	"os"
	"os/exec"
	"path/filepath"
)

var serviceName = "MYSQL"
var originalMysqlConfig = ""
var bakMysqlConfig = filepath.Join(agent_runtime.OutDir, "mysql.ini.bak")

func findMySQLServiceName() (string, error) {
	// Command to list all services, looking for MySQL
	//cmd := exec.Command("cmd", "/c", "sc", "query", "type=", "service", "state=", "all", "|", "findstr", "/i", "mysql")
	cmd := exec.Command("cmd", "/c", "net start | findstr /i MYSQL")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		utils.LogError("Failed to find MySQL service name: " + err.Error())
		return "", err
	}

	// Process `out` to find the MySQL service name
	// This is a simplified example. You'll need to parse the output to find the exact service name.
	return out.String(), nil
}

func platformRestartMysqlSkipGrantTables(mysqldCmd string) error {
	serviceName, err := findMySQLServiceName()
	if err != nil {
		return err
	}
	utils.LogInfo("Found MySQL service name: " + serviceName)
	mysqldProcess := utils.PlatformFindProcessAll("mysqld")
	cmdline, _ := mysqldProcess.Cmdline()

	//find config file from cmdline
	//mysqld --defaults-file="C:\ProgramData\MySQL\MySQL Server 8.0\my.ini" --console

	if mysqldProcess == nil {
		utils.LogError("mysqld not found")
		return errors.New("mysqld not found")
	}

	_, err = utils.RunCmd("net stop " + serviceName)

	if err != nil {
		utils.LogError("Failed to stop MySQL service: " + err.Error())
	}

	utils.LogInfo(cmdline)

	return err
}

func platformUseMysqldump(mysqldCmd string) (string, error) {
	mysqlDumpCmd, err := exec.LookPath("mysqldump.exe")
	if _, err := os.Stat(mysqlDumpCmd); os.IsNotExist(err) {
		mysqlDumpCmd = filepath.Join(filepath.Dir(mysqldCmd), "mysqldump.exe")
		if _, err := os.Stat(mysqlDumpCmd); os.IsNotExist(err) {
			mysqlDumpCmd = ""
			utils.LogError("mysqldump not found")
			return "", errors.New("mysqldump not found")
		}
	}
	file := sqlName("mysqldump")
	utils.LogInfo(mysqlDumpCmd)
	_, err = utils.RunCmd(fmt.Sprintf("%s --all-databases > %s", mysqlDumpCmd, file))
	if err != nil {
		utils.LogError(err.Error())
	}
	return file, err
}

func platformRestartMysqlDefault() (err error) {
	return
}
