//go:build linux

package utils

import (
	"errors"
	"go-agent/agent_runtime"
	"os/exec"
	"os/user"
	"strings"
)

func osInitBefore() error {

	currentUser, err := user.Current()
	if err != nil {
		return err
	}
	if currentUser.Uid != "0" {
		return errors.New("run it by root")
	}
	return nil
}

func osInitAfter() {

	// do something linux only
}

func RunCmd(cmd string) (string, error) {
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	return string(out), err
}

func IsAbsolutePath(path string) bool {
	if len(path) == 0 {
		return false
	}
	if path[0] == '/' {
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
			if strings.HasSuffix(exe, matchString) {
				env, _ := p.Environ()
				return exe, env

			}
		}
	}
	return "", nil
}
