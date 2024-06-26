//go:build linux

package utils

import (
	"errors"
	"github.com/elastic/go-sysinfo"
)

func osInitBefore() error {
	// do something linux only
	//assert run it by root
	process, err := sysinfo.Self()
	if err != nil {
		return err
	}
	user, err := process.User()
	if err != nil {
		return err
	}
	if user.UID != "0" {
		return errors.New("run it by root")
	}
	return nil
}

func osInitAfter() {

	// do something linux only
}
