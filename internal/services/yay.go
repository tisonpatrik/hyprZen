package services

import (
	"fmt"
	"log"
	"os"
)

func InstallYay() error {
	log.Println("InstallYay: Starting installation...")
	
	// Get the script path relative to the current working directory
	scriptPath := "scripts/fake_install.sh"
	
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

func InstallFake() error {
	log.Println("InstallFake: Starting fake installation...")
	
	// Get the script path relative to the current working directory
	scriptPath := "scripts/fake_install.sh"
	
	// Check if script exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		log.Printf("InstallFake: Script not found at %s", scriptPath)
		return fmt.Errorf("fake install script not found: %w", err)
	}
	
	log.Printf("InstallFake: Executing script: %s", scriptPath)
	
	// Execute the script
	if err := runCommand(scriptPath); err != nil {
		log.Printf("InstallFake: Script execution failed: %v", err)
		return fmt.Errorf("failed to execute fake install script: %w", err)
	}
	
	log.Println("InstallFake: Fake installation completed successfully")
	return nil
}