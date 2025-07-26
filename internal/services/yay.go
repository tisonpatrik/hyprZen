package services

import (
	"fmt"
	"log"
)

func InstallYay() error {
	log.Println("InstallYay: Starting installation...")
	
	if err := pacmanInstall("base-devel"); err != nil {
		log.Printf("InstallYay: Failed to install base-devel: %v", err)
		return fmt.Errorf("failed to install base-devel: %w", err)
	}
	log.Println("InstallYay: base-devel installed successfully")

	// Check if yay is already installed using yay --version
	if err := runCommandString("yay --version"); err == nil {
		log.Println("InstallYay: yay is already installed")
		return nil // yay is already installed
	}
	log.Println("InstallYay: yay not found, proceeding with installation")

	// Use /tmp directory like the original script
	tempDir := "/tmp/yay-bin"
	
	// Clone yay-bin in /tmp
	log.Println("InstallYay: Cloning yay-bin...")
	if err := runCommandString("cd /tmp && git clone https://aur.archlinux.org/yay-bin.git"); err != nil {
		log.Printf("InstallYay: Failed to clone yay-bin: %v", err)
		return fmt.Errorf("failed to clone yay-bin: %w", err)
	}
	log.Println("InstallYay: yay-bin cloned successfully")

	// Build and install yay
	log.Println("InstallYay: Building and installing yay...")
	if err := runCommandString(fmt.Sprintf("cd %s && makepkg -si --noconfirm", tempDir)); err != nil {
		log.Printf("InstallYay: Failed to build and install yay: %v", err)
		return fmt.Errorf("failed to build and install yay: %w", err)
	}
	log.Println("InstallYay: yay built and installed successfully")

	// Clean up
	log.Println("InstallYay: Cleaning up...")
	if err := runCommandString("rm -rf /tmp/yay-bin"); err != nil {
		log.Printf("InstallYay: Failed to clean up yay-bin directory: %v", err)
		return fmt.Errorf("failed to clean up yay-bin directory: %w", err)
	}
	log.Println("InstallYay: Cleanup completed")

	// Add fun and color to the pacman installer
	log.Println("InstallYay: Updating pacman.conf...")
	if err := runCommandString("sudo sed -i '/^\\[options\\]/a Color\\nILoveCandy' /etc/pacman.conf"); err != nil {
		log.Printf("InstallYay: Failed to update pacman.conf: %v", err)
		return fmt.Errorf("failed to update pacman.conf: %w", err)
	}
	log.Println("InstallYay: pacman.conf updated")

	// Verify yay is working
	log.Println("InstallYay: Verifying yay installation...")
	if err := runCommandString("yay --version"); err != nil {
		log.Printf("InstallYay: yay installation verification failed: %v", err)
		return fmt.Errorf("yay installation verification failed: %w", err)
	}
	log.Println("InstallYay: yay installation verified successfully")

	return nil
}