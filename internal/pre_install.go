package internal

import (
	"fmt"
)

func PreInstallSetup() error {
	if err := updateAndUpgrade(); err != nil {
		return err
	}

	if err := setupFlatpak(); err != nil {
		return err
	}

	return nil
}

func updateAndUpgrade() error {

	if err := runCommand("sudo", "zypper", "refresh"); err != nil {
		return fmt.Errorf("zypper refresh failed: %w", err)
	}

	if err := runCommand(
		"sudo", "zypper",
		"--non-interactive",
		"dup",
	); err != nil {
		return fmt.Errorf("zypper dup failed: %w", err)
	}

	return nil
}

func setupFlatpak() error {
	fmt.Println("📦 Installing Flatpak if not present...")
	if err := runCommand("sudo", "zypper", "--non-interactive", "--auto-agree-with-licenses", "in", "flatpak"); err != nil {
		return fmt.Errorf("failed to install flatpak: %w", err)
	}

	fmt.Println("🌍 Adding Flathub as user remote if it doesn't exist...")
	err := runCommand(
		"flatpak", "--user", "remote-add", "--if-not-exists",
		"flathub", "https://flathub.org/repo/flathub.flatpakrepo",
	)
	if err != nil {
		return fmt.Errorf("failed to add Flathub repo: %w", err)
	}

	fmt.Println("✅ Flatpak setup complete (user mode).")
	return nil
}
