package database

import (
	"bufio"
	"errors"
	"fmt"
	"go-agent/utils"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 其实不生效,会被覆盖
var originalMysqlConfig = "/etc/mysql/my.cnf"
var bakMysqlConfig = "/etc/mysql/cnf.bak"

func platformRestartMysqlSkipGrantTables(_ string) (err error) {

	//在 /etc/mysql/ 目录下找到包含[mysqld]的配置文件

	_ = os.Remove(bakMysqlConfig)
	originalMysqlConfig, err = findMysqldConfig("/etc/mysql")
	if err != nil {
		utils.LogError(fmt.Sprintf("Failed to find [mysqld] config in /etc/mysql: %s", err))
		return err
	}
	utils.LogInfo(fmt.Sprintf("Found [mysqld] config in %s", originalMysqlConfig))

	err = copyBakAndReplaceWithSkipGrantTables(originalMysqlConfig, bakMysqlConfig)
	if err != nil {
		return err
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

func platformUseMysqldump(mysqldCmd string) (string, error) {
	mysqlDumpCmd, err := exec.LookPath("mysqldump")
	if _, err := os.Stat(mysqlDumpCmd); os.IsNotExist(err) {
		mysqlDumpCmd = filepath.Join(filepath.Dir(mysqldCmd), "mysqldump")
		if _, err := os.Stat(mysqlDumpCmd); os.IsNotExist(err) {
			mysqlDumpCmd = ""
			utils.LogError("mysqldump not found")
			return "", errors.New("mysqldump not found")
		}
	}
	file := sqlName("mysqldump")
	_, err = utils.RunCmd(fmt.Sprintf("%s --all-databases > %s", mysqlDumpCmd, file))
	if err != nil {
		utils.LogError(err.Error())
	}
	return file, err
}

func platformRestartMysqlDefault() (err error) {
	if linkTargetPath == "" {
		err = utils.MoveFile(bakMysqlConfig, originalMysqlConfig)
	} else {
		err = os.Remove(originalMysqlConfig)
		err := os.Link(linkTargetPath, originalMysqlConfig)
		if err != nil {
			utils.LogError(fmt.Sprintf("Failed to link %s to %s: %s", linkTargetPath, originalMysqlConfig, err))
		}
	}
	_, err = utils.RunCmd("systemctl restart mysql")
	if err != nil {
		utils.LogError(fmt.Sprintf("Failed to start MySQL with default: %s", err))
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
