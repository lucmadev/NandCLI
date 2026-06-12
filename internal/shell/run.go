package shell

import (
	"os/exec"
)

func Run(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)

	return cmd.CombinedOutput()

}
