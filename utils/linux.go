//go:build linux

package utils

import (
	"errors"
	"os/exec"
	"os/user"
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
	out, err := exec.Command("bash", "-c", cmd).Output()
	return string(out), err
}
