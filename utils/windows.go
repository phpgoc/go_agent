//go:build windows

package utils

import (
	"errors"
	"go-agent/agent_runtime"
	"go-agent/utils/windows"
	"os/exec"
	"strings"
)

func osInitBefore() error {
	// do something Windows only
	isAdmin, err := windows.CurrentProcessRunAsAdmin()
	if err != nil {
		return err
	}
	if !isAdmin {
		return errors.New("run it by administrator")
	}
	return nil
}

func osInitAfter() {
	// do something windows only

}

func RunCmd(cmd string) (string, error) {
	out, err := exec.Command("cmd", "/C", cmd).CombinedOutput()
	return string(out), err
}

func IsAbsolutePath(path string) bool {

	if strings.Contains(path, ":") {
		return true
	}
	return false
}

func platformFindInProcess(matchStringArray []string) string {
	for _, p := range agent_runtime.GetProcesses() {
		exe, err := p.Exe()
		if err != nil {
			continue
		}
		for _, matchString := range matchStringArray {
			if strings.HasSuffix(exe, matchString+".exe") {
				return exe
			}
		}
	}
	return ""
}
