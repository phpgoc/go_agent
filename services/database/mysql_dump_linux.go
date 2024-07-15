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

// 其实不生效,会被覆盖
var originalMysqlConfig = "/etc/mysql/my.cnf"
var bakMysqlConfig = "/etc/mysql/cnf.bak"
var linkTargetPath string

func platformRestartMysqlSkipGrantTables() error {

	//在 /etc/mysql/ 目录下找到包含[mysqld]的配置文件

	_ = os.Remove(bakMysqlConfig)

	originalMysqlConfig, err := findMysqldConfig("/etc/mysql")
	if err != nil {
		utils.LogError(fmt.Sprintf("Failed to find [mysqld] config in /etc/mysql: %s", err))
		return err

	}
	utils.LogInfo(fmt.Sprintf("Found [mysqld] config in %s", originalMysqlConfig))
	sourceFileContent, err := utils.ReadFile(originalMysqlConfig)

	//判断是否是软链接
	linkTargetPath, err = os.Readlink(originalMysqlConfig)
	if err != nil {
		//copy
		_, err = utils.RunCmd(fmt.Sprintf("cp %s %s", originalMysqlConfig, bakMysqlConfig))
	} else {
		//unlink
		_, err = utils.RunCmd(fmt.Sprintf("rm %s", originalMysqlConfig))
	}

	newFile, _ := os.OpenFile(originalMysqlConfig, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	defer func(newFile *os.File) {
		err := newFile.Close()
		if err != nil {

		}
	}(newFile)

	foundMysqld := false
	updatedBindAddress := false
	updatePort := false

	for _, line := range strings.Split(sourceFileContent, "\n") {

		// Check if we're in the [mysqld] section
		if strings.TrimSpace(line) == "[mysqld]" {
			foundMysqld = true
			line += "\n skip-grant-tables"
		} else if foundMysqld && strings.Contains(line, "bind-address") {
			line = "bind-address = 127.0.0.1"
			updatedBindAddress = true
		} else if foundMysqld && strings.Contains(line, "port") {
			line = "port = 3306"
			updatePort = true
		}

		_, _ = newFile.WriteString(line + "\n")

	}

	// If we didn't find the bind-address line, add it under the [mysqld] section
	if foundMysqld && !updatedBindAddress {
		_, err := newFile.WriteString("bind-address =127.0.0.1:3306\n")
		if err != nil {
			return err
		}
	}

	if foundMysqld && !updatePort {
		_, err := newFile.WriteString("port = 3306\n")
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
	utils.LogWarn("MySQL restarted with --skip-grant-tables")
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
	utils.LogWarn("MySQL restarted with default configuration")
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
			_ = f.Close()
			continue // Skip files that cannot be opened
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {

			if strings.HasPrefix(strings.TrimSpace(scanner.Text()), "[mysqld]") {
				return filePath, nil // Found the config file
			}
		}
		_ = f.Close()
	}

	return "", fmt.Errorf("no [mysqld] config found in %s", dirPath)
}
