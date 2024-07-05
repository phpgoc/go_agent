package apache

func platformReadEnvFile(httpDefaultRoot string, envMap map[string]string) map[string]string {
	//windows下不需要读取envvars,直接返回envMap
	return envMap
}
