package monitor

import (
	"fmt"
	"os/exec"
)

func GetReport() string {
	cpu, _ := exec.Command("sh", "-c", "top -bn1 | grep 'Cpu'").CombinedOutput()
	mem, _ := exec.Command("sh", "-c", "free -h").CombinedOutput()
	df, _ := exec.Command("sh", "-c", "df -h /").CombinedOutput()

	return fmt.Sprintf(
		"--- RECURSOS ---\nCPU: %s\nMEMORIA:\n%s\nDISCO:\n%s\n",
		string(cpu), string(mem), string(df),
	)
}
