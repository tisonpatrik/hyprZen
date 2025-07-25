package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}



func zypperInstallMany(pkgs []string) error {
	if len(pkgs) == 0 {
		return nil
	}
	args := append([]string{"zypper", "in"}, pkgs...)
	if err := runCommand("sudo", args...); err != nil {
		return fmt.Errorf("failed to install packages %v: %w", pkgs, err)
	}
	return nil
}
