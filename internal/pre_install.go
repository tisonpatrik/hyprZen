package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func PreInstallSetup() error {
	if err := initSudoSession(); err != nil {
		return err
	}
	if err := updateAndUpgrade(); err != nil {
		return err
	}
	if err := setupFlatpak(); err != nil {
		return err
	}
	if err := installOpi(); err != nil {
		return err
	}
	if err := installCodecs(); err != nil {
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
	if err := runCommand("sudo", "zypper", "--non-interactive", "in", "flatpak"); err != nil {
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

func installOpi() error {
	if err := runCommand("sudo", "zypper", "--non-interactive", "in", "opi"); err != nil {
		return fmt.Errorf("failed to install opi: %w", err)
	}
	return nil
}

func installCodecs() error {
	fmt.Println("📦 Installing codecs via opi...")

	cmd := exec.Command("bash", "-c", "opi codecs")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("ℹ️  You may be prompted by opi for input — please follow instructions.")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run opi codecs: %w", err)
	}

	return nil
}

func initSudoSession() error {
	fmt.Println("🔐 Asking for sudo password to initialize session...")

	// Pro jistotu zrušme předchozí session (vynucení zadání hesla)
	exec.Command("sudo", "-K").Run()

	cmd := exec.Command("sudo", "-v")
	cmd.Stdin = os.Stdin   // aby prompt fungoval
	cmd.Stdout = os.Stdout // pro výpis dotazu
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to authenticate with sudo: %w", err)
	}

	return nil
}
