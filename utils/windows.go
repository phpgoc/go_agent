//go:build windows

package utils

import (
	"errors"
	"golang.org/x/sys/windows"
)

func osInitBefore() error {
	// do something Windows only
	isAdmin, err := isAdmin()
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

func isAdmin() (bool, error) {
	var sid *windows.SID

	// Well known SID for the Administrators group
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid)
	if err != nil {
		return false, err
	}

	defer windows.FreeSid(sid)

	token, err := windows.OpenCurrentProcessToken()
	if err != nil {
		return false, err
	}

	defer token.Close()

	groups, err := token.GetTokenGroups()
	if err != nil {
		return false, err
	}

	for _, group := range groups.AllGroups() {

		if windows.EqualSid(group.Sid, sid) {
			return true, nil
		}
	}

	return false, nil
}
