package services

import (
	"fmt"
	"log"
	"os"
)

func InstallYay() error {
	log.Println("InstallYay: Starting installation...")
	
	// Get the script path relative to the current working directory
	scriptPath := "scripts/install_yay.sh"
	
	// Check if script exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		log.Printf("InstallYay: Script not found at %s", scriptPath)
		return fmt.Errorf("install script not found: %w", err)
	}
	
	log.Printf("InstallYay: Executing script: %s", scriptPath)
	
	// Execute the script
	if err := runCommand(scriptPath); err != nil {
		log.Printf("InstallYay: Script execution failed: %v", err)
		return fmt.Errorf("failed to execute install script: %w", err)
	}
	
	log.Println("InstallYay: Installation completed successfully")
	return nil
}