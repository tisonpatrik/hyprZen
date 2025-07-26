package services

import (
	"fmt"
	"math/rand"
	"time"
)

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

// Install performs the complete HyprZen installation (dummy version)
func (i *InstallerService) Install() error {
	// Simulate installation steps with delays
	time.Sleep(2 * time.Second) // PreInstallSetup

	time.Sleep(3 * time.Second) // InstallSystem

	// Simulate error if requested
	if i.simulateError {
		return fmt.Errorf("network connection error during system installation")
	}

	time.Sleep(2 * time.Second) // InstallAps
	time.Sleep(1 * time.Second) // AddConfigs

	return nil
}

// GetPackages returns the list of packages to install
func (i *InstallerService) GetPackages() []string {
	packages := []string{
		"vegeutils",
		"libgardening",
		"currykit",
		"spicerack",
		"fullenglish",
		"eggy",
		"bad-kitty",
		"chai",
		"hojicha",
		"libtacos",
		"babys-monads",
		"libpurring",
		"currywurst-devel",
		"xmodmeow",
		"licorice-utils",
		"cashew-apple",
		"rock-lobster",
		"standmixer",
		"coffee-CUPS",
		"libesszet",
		"zeichenorientierte-benutzerschnittstellen",
		"schnurrkit",
		"old-socks-devel",
		"jalapeño",
		"molasses-utils",
		"xkohlrabi",
		"party-gherkin",
		"snow-peas",
		"libyuzu",
	}

	// Shuffle and add version numbers like in the example
	pkgs := make([]string, len(packages))
	copy(pkgs, packages)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(pkgs), func(i, j int) {
		pkgs[i], pkgs[j] = pkgs[j], pkgs[i]
	})

	for k := range pkgs {
		pkgs[k] += fmt.Sprintf("-%d.%d.%d", rand.Intn(10), rand.Intn(10), rand.Intn(10))
	}

	return pkgs
}
