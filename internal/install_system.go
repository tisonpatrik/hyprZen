package internal

import "fmt"

func InstallSystem() error {
	packages := []string{"hyprland-devel", "hyprland-qtutils", "hyprland"}
	if err := zypperInstallMany(packages); err != nil {
		return fmt.Errorf("failed to install hyprland: %w", err)
	}
	return nil
}
