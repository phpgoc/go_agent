package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func WaitUntil(cmd, expect string, interval, timeout uint) error {
	var i uint
	for i = 0; timeout != 0 && i < timeout/interval; i += interval {
		out, err := RunCmd(cmd)
		if err != nil {
			return err
		}
		// contain
		if strings.Contains(out, expect) {
			return nil
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
	return nil

}

// FindCommandFromPathAndProcessByMatchStringArray matchStringArray 具有优先级，找到第一个就返回
func FindCommandFromPathAndProcessByMatchStringArray(matchStringArray []string) (cmd string, env []string) {

	//先找process
	cmd, env = platformFindInProcess(matchStringArray)
	if cmd != "" {
		return
	}
	//trim every input string
	for i, s := range matchStringArray {
		matchStringArray[i] = strings.TrimSpace(s)
	}
	//look path
	for _, p := range matchStringArray {
		absPath, err := exec.LookPath(p)
		if err == nil {
			return absPath, nil
		}
	}

	return "", nil
}

func GetSystemEnvVars() map[string]string {
	envVars := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			envVars[pair[0]] = pair[1]
		}
	}
	return envVars
}

func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer func(sourceFile *os.File) {
		err := sourceFile.Close()
		if err != nil {
			LogError(fmt.Sprintf("failed to close source file: %s", err))
		}
	}(sourceFile)

	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer func(destFile *os.File) {
		err := destFile.Close()
		if err != nil {
			LogError(fmt.Sprintf("failed to close destination file: %s", err))
		}
	}(destFile)

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}

func MoveFile(src, dst string) error {
	// Ensure the source file exists.
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return err
	}

	// Ensure the destination directory exists.
	if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
		return err
	}

	// Open the source file.
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {
			LogError(fmt.Sprintf("failed to close source file: %s", err))
		}
	}(source)

	// Create the destination file.
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(destination *os.File) {
		err := destination.Close()
		if err != nil {
			LogError(fmt.Sprintf("failed to close destination file: %s", err))
		}
	}(destination)

	// Copy the source file to the destination.
	if _, err := io.Copy(destination, source); err != nil {
		return err
	}

	// Delete the original source file.
	return os.Remove(src)
}
