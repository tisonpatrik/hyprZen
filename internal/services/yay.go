package services

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallYay() error {
	// Check if yay is already installed
	if _, err := exec.LookPath("yay"); err == nil {
		return nil // yay is already installed
	}

	// Use a temporary directory that gets cleaned up automatically
	tempDir, err := os.MkdirTemp("", "yay-install-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir) // Clean up on exit

	// Install base-devel if needed
	if err := pacmanInstall("base-devel"); err != nil {
		return fmt.Errorf("failed to install base-devel: %w", err)
	}

	// Clone and build yay in temp directory
	if err := runCommandString(fmt.Sprintf("git clone https://aur.archlinux.org/yay-bin.git %s", tempDir)); err != nil {
		return fmt.Errorf("failed to clone yay-bin: %w", err)
	}

	// Build and install yay
	if err := runCommandString(fmt.Sprintf("cd %s && makepkg -si --noconfirm", tempDir)); err != nil {
		return fmt.Errorf("failed to build and install yay: %w", err)
	}

	// Verify yay is working
	if err := runCommandString("yay --version"); err != nil {
		return fmt.Errorf("yay installation verification failed: %w", err)
	}

	return nil
}