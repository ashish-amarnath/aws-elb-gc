package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/golang/glog"
)

// RunBashCmd runs a supplied bash command
func RunBashCmd(cmd string) (res string, err error) {
	glog.V(8).Infof("Running [%s]\n", cmd)
	toRun := exec.Command("bash", "-c", cmd)
	var stderr bytes.Buffer
	toRun.Stderr = &stderr
	out, err := toRun.Output()
	if err != nil {
		res = ""
		errMsg := fmt.Sprintf("ERROR: %s: %v", stderr.String(), err)
		err = fmt.Errorf(errMsg)
		glog.Error(errMsg)
	}
	res = strings.TrimSpace(string(out))
	return
}
