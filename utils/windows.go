//go:build windows

package utils

import (
	"errors"
	"go-agent/utils/windows"
	"os/exec"
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
	out, err := exec.Command("cmd", "/C", cmd).Output()
	return string(out), err
}
