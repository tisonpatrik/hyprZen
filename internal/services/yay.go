package services

import (
	"fmt"
)

func InstallYay() error {
	if err := pacmanInstall("base-devel"); err != nil {
		return fmt.Errorf("failed to install base-devel: %w", err)
	}

	// Check if yay is already installed using yay --version
	if err := runCommandString("yay --version"); err == nil {
		return nil // yay is already installed
	}

	// Use /tmp directory like the original script
	tempDir := "/tmp/yay-bin"
	
	// Clone yay-bin in /tmp
	if err := runCommandString("cd /tmp && git clone https://aur.archlinux.org/yay-bin.git"); err != nil {
		return fmt.Errorf("failed to clone yay-bin: %w", err)
	}

	// Build and install yay
	if err := runCommandString(fmt.Sprintf("cd %s && makepkg -si --noconfirm", tempDir)); err != nil {
		return fmt.Errorf("failed to build and install yay: %w", err)
	}

	// Clean up
	if err := runCommandString("rm -rf /tmp/yay-bin"); err != nil {
		return fmt.Errorf("failed to clean up yay-bin directory: %w", err)
	}

	// Add fun and color to the pacman installer
	if err := runCommandString("sudo sed -i '/^\\[options\\]/a Color\\nILoveCandy' /etc/pacman.conf"); err != nil {
		return fmt.Errorf("failed to update pacman.conf: %w", err)
	}

	// Verify yay is working
	if err := runCommandString("yay --version"); err != nil {
		return fmt.Errorf("yay installation verification failed: %w", err)
	}

	return nil
}