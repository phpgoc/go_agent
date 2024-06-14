// +build:linux
package linux

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func RunCmd(cmd string) (string, error) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// timeout if 0 then wait forever
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

var LogFileName = "/tmp/xxx/agent.log"
var LogFile *os.File

func WriteLogFile(log string, level string) (err error) {
	// log to file
	_, err = LogFile.WriteString(fmt.Sprintf("%s %s\n", level, log))
	return err
}
