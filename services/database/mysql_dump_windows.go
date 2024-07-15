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
	"regexp"
	"strings"
)

var serviceName = ""
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
	return strings.TrimSpace(out.String()), nil
}

func platformRestartMysqlSkipGrantTables(mysqldCmd string) (err error) {
	serviceName, err = findMySQLServiceName()
	if err != nil {
		return err
	}
	utils.LogInfo("Found MySQL service name: " + serviceName)
	mysqldProcess := utils.PlatformFindProcessAll("mysqld")
	cmdline, _ := mysqldProcess.Cmdline()

	if mysqldProcess == nil {
		utils.LogError("mysqld not found")
		return errors.New("mysqld not found")
	}

	//find config file from cmdline
	//mysqld --defaults-file="C:\ProgramData\MySQL\MySQL Server 8.0\my.ini" --console
	re := regexp.MustCompile(`--defaults-file="(.+?)"`)
	matches := re.FindStringSubmatch(cmdline)
	utils.LogInfo(cmdline)
	if len(matches) < 2 {
		re = regexp.MustCompile(`"--defaults-file=(.+?)"`)
		matches = re.FindStringSubmatch(cmdline)
		if len(matches) < 2 {
			utils.LogError("Failed to find config file from cmdline")
			return errors.New("Failed to find config file from cmdline")
		}
	}
	originalMysqlConfig = matches[1]
	utils.LogInfo("Found MySQL config file: " + originalMysqlConfig)
	copyBakAndReplaceWithSkipGrantTables(originalMysqlConfig, bakMysqlConfig)

	_, err = utils.RunCmd("net stop " + serviceName)

	if err != nil {
		utils.LogError("Failed to stop MySQL service : " + err.Error())
	}
	_, err = utils.RunCmd("net start " + serviceName)

	if err != nil {
		utils.LogError("Failed to start MySQL service with skip grant tables : " + err.Error())
	}

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
	os.MkdirAll(filepath.Dir(file), os.ModePerm)
	outputFile, err := os.Create(file)
	if err != nil {
		utils.LogError(fmt.Sprintf("Failed to create output file: %s", err))
		return "", err
	}
	defer outputFile.Close()

	// Prepare the mysqldump command without using shell redirection
	cmd := exec.Command(mysqlDumpCmd, "--all-databases")

	// Set the output file as the standard output of the command
	cmd.Stdout = outputFile

	// Execute the command
	err = cmd.Run()
	if err != nil {
		utils.LogError(fmt.Sprintf("Failed to execute mysqldump command: %s", err))
		return "", err
	}

	utils.LogWarn(fmt.Sprintf(`Executed: "%s" --all-databases > "%s"`, mysqlDumpCmd, file))
	return file, nil
}

func platformRestartMysqlDefault() (err error) {
	if linkTargetPath == "" {
		err = utils.MoveFile(bakMysqlConfig, originalMysqlConfig)
	} else {
		err = os.Remove(originalMysqlConfig)
		err = os.Link(linkTargetPath, originalMysqlConfig)
		if err != nil {
			utils.LogError(fmt.Sprintf("Failed to link %s to %s: %s", linkTargetPath, originalMysqlConfig, err))
		}
	}
	_, err = utils.RunCmd("net stop " + serviceName)
	if err != nil {
		utils.LogError(fmt.Sprintf("Failed to stop %s with skip grant tables: %s", serviceName, err))
		return err
	}
	_, err = utils.RunCmd("net start " + serviceName)
	if err != nil {
		utils.LogError(fmt.Sprintf("Failed to start %s with default: %s", serviceName, err))
		return err
	}

	utils.LogWarn("MySQL restarted with default configuration")

	return nil
}
