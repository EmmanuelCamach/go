package exec

import (
	"os/exec"
)

func Run(cmd string) string {
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return string(out) + "\n" + err.Error()
	}
	return string(out)
}
