//go:build windows

package windows

import "os/exec"

func RunCmd(cmd string) (string, error) {
	out, err := exec.Command("cmd", "/C", cmd).Output()
	return string(out), err
}
