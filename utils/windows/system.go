//go:build windows

package windows

import (
	golang_windows "golang.org/x/sys/windows"
)

func CurrentProcessRunAsAdmin() (bool, error) {
	var sid *golang_windows.SID

	// Well known SID for the Administrators group
	err := golang_windows.AllocateAndInitializeSid(
		&golang_windows.SECURITY_NT_AUTHORITY,
		2,
		golang_windows.SECURITY_BUILTIN_DOMAIN_RID,
		golang_windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid)
	if err != nil {
		return false, err
	}

	defer golang_windows.FreeSid(sid)

	token, err := golang_windows.OpenCurrentProcessToken()
	if err != nil {
		return false, err
	}

	defer token.Close()

	groups, err := token.GetTokenGroups()
	if err != nil {
		return false, err
	}

	for _, group := range groups.AllGroups() {

		if golang_windows.EqualSid(group.Sid, sid) {
			return true, nil
		}
	}
	return false, nil
}

func IsAdmin(userId string) (bool, error) {
	// Convert the user ID to a SID
	//sid, err := golang_windows.StringToSid(userId)
	sid, err := golang_windows.StringToSid(userId)
	if err != nil {
		return false, err
	}

	// Well known SID for the Administrators group
	//SID S-1-5-32-544 是Windows系统中预定义的，用于表示"Administrators"组的SID。这个SID在所有的Windows系统中都是相同的。因此，你可以安全地在代码中使用这个SID来检查一个用户是否是管理员。
	adminSid, err := golang_windows.StringToSid("S-1-5-32-544")
	if err != nil {
		return false, err
	}

	// Check if the user SID is equal to the Administrators group SID
	isAdmin := golang_windows.EqualSid(sid, adminSid)

	return isAdmin, nil
}
