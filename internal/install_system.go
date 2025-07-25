package internal

import "fmt"

var (

	hyprlandia = []string{
		"hyprland-devel",
		"hyprland-qtutils",
		"hyprland",
		"hyprlock",
		"hyprshot",
		"hypridle",
		"hyprland-wallpapers",
		"xdg-desktop-portal-hyprland",
	}
	systemTools = []string{
		"blueman",
		"xdg-desktop-portal-gtk",
		"gtk4-devel",
		"brightnessctl",
		"playerctl",
		"polkit-gnome",
		"libqalculate",
	}
	uiTools = []string{
		"waybar",
		"mako",
		"swaybg",
	}
	opiPackages = []string{
		"uwsm",
	}
)

func InstallSystem() error {
	if err := installHyprland(); err != nil {
		return err
	}
	if err := installTools(); err != nil {
		return err
	}
	if err := opiInstallMany(opiPackages); err != nil {
		return err
	}
	return nil
}

func installHyprland() error {
	if err := zypperInstallMany(hyprlandia); err != nil {
		return fmt.Errorf("failed to install hyprland: %w", err)
	}
	return nil
}

func installTools() error {
	all := append(uiTools, systemTools...)
	if err := zypperInstallMany(all); err != nil {
		return fmt.Errorf("failed to install tools: %w", err)
	}
	return nil
}
