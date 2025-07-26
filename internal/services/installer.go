package services

// InstallerService handles all installation-related operations
type InstallerService struct {
	simulateError bool
}

// NewInstallerService creates a new installer service instance
func NewInstallerService() *InstallerService {
	return &InstallerService{
		simulateError: false,
	}
}

// NewInstallerServiceWithError creates a new installer service that simulates errors
func NewInstallerServiceWithError() *InstallerService {
	return &InstallerService{
		simulateError: true,
	}
}

// InstallStep represents a single installation step
type InstallStep struct {
	Name     string
	Packages []string
	Action   func() error
}

// Install performs the complete HyprZen installation and returns the list of installation steps
func (i *InstallerService) Install() []InstallStep {
	steps := []InstallStep{
		{
			Name: "Installing yay package manager",
			Action: func() error {
				return InstallYay()
			},
		},
	}
	
	return steps
}

