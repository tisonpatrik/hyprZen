package services

import (
	"fmt"
	"log"
)

func InstallYay() error {
	log.Println("InstallYay: Starting installation...")
	
	// First, ensure git is available
	log.Println("InstallYay: Checking for git...")
	if err := runCommandString("git --version"); err != nil {
		log.Println("InstallYay: Git not found, installing git...")
		if err := pacmanInstall("git"); err != nil {
			log.Printf("InstallYay: Failed to install git: %v", err)
			return fmt.Errorf("failed to install git: %w", err)
		}
		log.Println("InstallYay: Git installed successfully")
	} else {
		log.Println("InstallYay: Git is already available")
	}
	
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
	// First remove any existing directory
	if err := runCommandString("rm -rf /tmp/yay-bin"); err != nil {
		log.Printf("InstallYay: Failed to clean existing directory: %v", err)
	}
	
	// Clone using git command directly
	if err := runCommand("git", "clone", "https://aur.archlinux.org/yay-bin.git", "/tmp/yay-bin"); err != nil {
		log.Printf("InstallYay: Failed to clone yay-bin: %v", err)
		// Let's also check if git is available
		if gitErr := runCommandString("git --version"); gitErr != nil {
			log.Printf("InstallYay: Git is not available: %v", gitErr)
			return fmt.Errorf("git is not available: %w", gitErr)
		}
		return fmt.Errorf("failed to clone yay-bin: %w", err)
	}
	log.Println("InstallYay: yay-bin cloned successfully")

	// Build and install yay
	log.Println("InstallYay: Building and installing yay...")
	// Change to the temp directory and run makepkg
	if err := runCommand("sh", "-c", fmt.Sprintf("cd %s && makepkg -si --noconfirm", tempDir)); err != nil {
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