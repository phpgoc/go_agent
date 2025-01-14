//go:build windows

package utils

import (
	"errors"
	"github.com/shirou/gopsutil/v4/process"
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

func platformFindInProcess(matchStringArray []string) (cmd string, env []string) {
	for _, p := range agent_runtime.GetProcesses() {
		exe, err := p.Exe()
		if err != nil {
			continue
		}

		for _, matchString := range matchStringArray {
			if strings.HasSuffix(exe, matchString+".exe") {
				env, _ := p.Environ()
				return exe, env
			}
		}
	}
	return "", nil
}

func PlatformFindProcessAll(match string) *process.Process {
	match = match + ".exe"
	for _, p := range agent_runtime.GetProcesses() {
		exe, err := p.Exe()
		if err != nil {
			continue
		}
		if strings.HasSuffix(exe, match) {
			return p
		}
	}
	return nil
}
