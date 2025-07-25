package internal

import "fmt"

func InstallSystem() error {

	if err := installHyprland(); err != nil {
		return err
	}
	if err := installDrivers(); err != nil {
		return err
	}
	return nil
}

func installHyprland() error {

	packages := []string{"hyprland-devel", "hyprland-qtutils", "hyprland", "hyprlock", "hyprshot", "waybar", "hyprland-wallpapers"}
	if err := zypperInstallMany(packages); err != nil {
		return fmt.Errorf("failed to install hyprland: %w", err)
	}
	return nil
}

func installDrivers() error {
	packages := []string{"blueman", "xdg-desktop-portal-hyprland", "xdg-desktop-portal-gtk"}
	if err := zypperInstallMany(packages); err != nil {
		return fmt.Errorf("failed to install drivers: %w", err)
	}
	return nil
}


