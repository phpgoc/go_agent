package utils

import (
	"os"
	"os/exec"
	"strings"
	"time"
)

func WaitUntil(cmd, expect string, interval, timeout uint) error {
	var i uint
	for i = 0; timeout != 0 && i < timeout/interval; i += interval {
		out, err := RunCmd(cmd)
		if err != nil {
			return err
		}
		// contain
		if strings.Contains(out, expect) {
			return nil
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
	return nil

}

// FindCommandFromPathAndProcessByMatchStringArray matchStringArray 具有优先级，找到第一个就返回
func FindCommandFromPathAndProcessByMatchStringArray(matchStringArray []string) (cmd string, env []string) {

	//先找process
	cmd, env = platformFindInProcess(matchStringArray)
	if cmd != "" {
		return
	}
	//trim every input string
	for i, s := range matchStringArray {
		matchStringArray[i] = strings.TrimSpace(s)
	}
	//look path
	for _, p := range matchStringArray {
		absPath, err := exec.LookPath(p)
		if err == nil {
			return absPath, nil
		}
	}

	return "", nil
}

func GetSystemEnvVars() map[string]string {
	envVars := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			envVars[pair[0]] = pair[1]
		}
	}
	return envVars
}
