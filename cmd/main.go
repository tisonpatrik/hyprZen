package main

import (
	"fmt"
	"hyprzen/internal"
	"log"
)

func main() {
	if err := internal.PreInstallSetup(); err != nil {
		log.Fatalf("PreInstallSetup failed: %v", err)
	}
	if err := internal.InstallSystem(); err != nil {
		log.Fatalf("InstallSystem failed: %v", err)
	}
	if err := internal.InstallAps(); err != nil {
		log.Fatalf("InstallAps failed: %v", err)
	}
	if err := internal.AddConfigs(); err != nil {
		log.Fatalf("AddConfigs failed: %v", err)
	}
	
	fmt.Println("✅ All installation steps completed successfully!")
}
