package apache

import (
	"go-agent/utils"
	"path/filepath"
)

func platformReadEnvFile(httpDefaultRoot string, envMap map[string]string) map[string]string {

	//env dict ,this file name by guess
	envContent := utils.GetFirstAndLogError(
		func() (string, error) {
			return utils.ReadFile(filepath.Join(httpDefaultRoot, "envvars"))
		})
	KVMap := utils.InterpretSourceExportToGoMap(envContent, envMap)
	return KVMap
}
