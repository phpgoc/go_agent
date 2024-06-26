//go:build windows

package windows

import (
	"os/exec"
	"syscall"
)

func RunCmd(cmd string) (string, error) {
	out, err := exec.Command("cmd", "/C", cmd).Output()
	return string(out), err
}

func RunMsi(cmd string) ([]byte, error) {
	return exec.Command("msiexec", cmd).Output()
}

func RunWmic(cmd string) ([]byte, error) {
	wmic := exec.Command("wmic")
	wmic.SysProcAttr = &syscall.SysProcAttr{CmdLine: "/C " + cmd}
	return wmic.Output()
}

func RunPowershell(cmd string) ([]byte, error) {
	return exec.Command("powershell", "/Command", cmd).Output()
}
