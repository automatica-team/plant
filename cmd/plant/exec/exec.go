package exec

import (
	"os"
	"os/exec"
)

var dir, _ = os.Getwd()

func Cd(v string) {
	dir = v
}

func Command(name string, arg ...string) error {
	return command(false, name, arg...)
}

func CommandSilent(name string, arg ...string) error {
	return command(true, name, arg...)
}

func command(silent bool, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Env = os.Environ()
	cmd.Dir = dir

	if !silent {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}
