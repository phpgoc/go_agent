package database

import (
	"bufio"
	"fmt"
	"go-agent/utils"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var originalMysqlConfig = "/etc/mysql/my.cnf"
var bakMysqlConfig = "/etc/mysql/my.cnf.bak"
var linkTargetPath string

func platformRestartMysqlSkipGrantTables() error {
	// Command to stop MySQL service. This might need to be adjusted based on your system.
	//_, err := utils.RunCmd("systemctl stop mysql")
	//
	//if err != nil {
	//	utils.LogError(fmt.Sprintf("Failed to stop MySQL service: %s", err))
	//	return err
	//}
	//在 /etc/mysql/ 目录下找到包含[mysqld]的配置文件

	originalMysqlConfig, err := findMysqldConfig("/etc/mysql")
	if err != nil {
		utils.LogError(fmt.Sprintf("Failed to find [mysqld] config in /etc/mysql: %s", err))
		return err

	}
	utils.LogInfo(fmt.Sprintf("Found [mysqld] config in %s", originalMysqlConfig))
	sourceFileContent, err := utils.ReadFile(originalMysqlConfig)

	linkTargetPath, err = os.Readlink(originalMysqlConfig)
	if err != nil {
		//copy
		_, err = utils.RunCmd(fmt.Sprintf("cp %s %s", originalMysqlConfig, bakMysqlConfig))
	} else {
		//unlink
		_, err = utils.RunCmd(fmt.Sprintf("rm %s", originalMysqlConfig))
	}

	newFile, _ := os.OpenFile("/etc/mysql/my.cnf", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	defer func(newFile *os.File) {
		err := newFile.Close()
		if err != nil {

		}
	}(newFile)

	// Variables to track whether we've found the [mysqld] section and updated the bind-address
	foundMysqld := false
	updatedBindAddress := false

	for _, line := range strings.Split(sourceFileContent, "\n") {

		// Check if we're in the [mysqld] section
		if strings.TrimSpace(line) == "[mysqld]" {
			foundMysqld = true
			line += "\n skip-grant-tables"
		} else if foundMysqld && strings.Contains(line, "bind-address") {
			line = "bind-address = 127.0.0.1"
			updatedBindAddress = true
		}
		_, _ = newFile.WriteString(line + "\n")

	}

	// If we didn't find the bind-address line, add it under the [mysqld] section
	if foundMysqld && !updatedBindAddress {
		_, err := newFile.WriteString("bind-address = 127.0.0.1\n")
		if err != nil {
			return err
		}
	}

	_, err = utils.RunCmd("systemctl restart mysql")
	if err != nil {
		utils.LogError(fmt.Sprintf("Failed to start MySQL with --skip-grant-tables: %s", err))
		return err
	}
	// 一直等到mysql启动
	// active比running更准确
	return utils.WaitUntil("systemctl status mysql", "active", 1, 0)

}

func platformRestartMysqlDefault() (err error) {
	if linkTargetPath == "" {
		_, err = utils.RunCmd(fmt.Sprintf("mv %s %s", bakMysqlConfig, originalMysqlConfig))
	} else {
		_, err = utils.RunCmd(fmt.Sprintf("rm %s", originalMysqlConfig))
		_, err = utils.RunCmd(fmt.Sprintf("ln -s %s %s", linkTargetPath, originalMysqlConfig))
	}
	_, err = utils.RunCmd("systemctl restart mysql")
	if err != nil {
		utils.LogError(fmt.Sprintf("Failed to start MySQL with --skip-grant-tables: %s", err))
		return err
	}
	// 一直等到mysql启动
	// active比running更准确
	return utils.WaitUntil("systemctl status mysql", "active", 1, 0)
}
func findMysqldConfig(dirPath string) (string, error) {
	var files []fs.DirEntry
	var err error

	if files, err = os.ReadDir(dirPath); err != nil {
		return "", err
	}

	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())
		if file.IsDir() {
			found, _ := findMysqldConfig(filePath)
			if found != "" {
				return found, nil
			}
		}

		f, err := os.Open(filePath)
		if err != nil {
			continue // Skip files that cannot be opened
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if scanner.Text() == "[mysqld]" {
				return filePath, nil // Found the config file
			}
		}
	}

	return "", fmt.Errorf("no [mysqld] config found in %s", dirPath)
}
