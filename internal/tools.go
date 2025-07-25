package internal

import "fmt"

func InstallTools() error {
	packages := []string{"ghostty", "bitwarden", "neovim"}
	if err := zypperInstallMany(packages); err != nil {
		return fmt.Errorf("failed to install hyprland: %w", err)
	}
	return nil
}
