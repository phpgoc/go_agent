//go:build windows

package windows

import (
	"os/exec"
	"syscall"
)

func RunWmic(cmd string) ([]byte, error) {
	wmic := exec.Command("wmic")
	wmic.SysProcAttr = &syscall.SysProcAttr{CmdLine: "/C " + cmd}
	return wmic.Output()
}

func RunPowershell(cmd string) ([]byte, error) {
	return exec.Command("powershell", "/Command", cmd).Output()
}
