package services

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runCommandString(command string) error {
	// Split the command string into command and arguments
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return fmt.Errorf("empty command")
	}
	
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func yayInstall(pkgs ...string) error {
	if len(pkgs) == 0 {
		return nil
	}
	args := append([]string{"yay", "-S", "--noconfirm", "--needed"}, pkgs...)
	if err := runCommand(args[0], args[1:]...); err != nil {
		return fmt.Errorf("failed to install packages %v: %w", pkgs, err)
	}
	return nil
}

func pacmanInstall(pkgs ...string) error {
	if len(pkgs) == 0 {
		return nil
	}
	args := append([]string{"sudo","pacman", "-S", "--noconfirm", "--needed"}, pkgs...)
	if err := runCommand(args[0], args[1:]...); err != nil {
		return fmt.Errorf("failed to install packages %v: %w", pkgs, err)
	}
	return nil
}
