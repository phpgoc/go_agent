package utils

import (
	"go-agent/runtime"
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
func FindCommandFromPathAndProcessByMatchStringArray(matchStringArray []string) string {
	//trim every input string
	for i, s := range matchStringArray {
		matchStringArray[i] = strings.TrimSpace(s)
	}
	//look path
	for _, p := range matchStringArray {
		_, err := exec.LookPath(p)
		if err == nil {
			return p
		}
	}

	for _, p := range runtime.Processes {
		exe, err := p.Exe()
		if err != nil {
			continue
		}
		for _, matchString := range matchStringArray {
			if strings.HasSuffix(exe, matchString) {
				return exe
			}
		}
	}
	return ""
}
