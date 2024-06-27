//go:build linux

package utils

import (
	"errors"
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
