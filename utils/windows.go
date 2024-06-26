//go:build windows

package utils

import (
	"errors"
	"go-agent/utils/windows"
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
