package internal

import (
	"fmt"
	"os/exec"
)

func PreInstallSetup() error {
	if err := initSudoSession(); err != nil {
		return err
	}
	if err := updateAndUpgrade(); err != nil {
		return err
	}
	if err := installDevelPattern(); err != nil{
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

func installDevelPattern() error {
        if err := runCommand("sudo", "zypper", "install", "-t", "pattern", "devel_basis"); err != nil {
                return fmt.Errorf("devel_basis pattern install failed: %w", err)
        }
        return nil
}

func updateAndUpgrade() error {

	if err := runCommand("sudo", "zypper", "refresh"); err != nil {
		return fmt.Errorf("zypper refresh failed: %w", err)
	}

	if err := runCommand(
		"sudo", "zypper",
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
	fmt.Println("🎵 Adding Packman repository for multimedia support...")

	err := runCommand(
		"sudo", "zypper", "ar", "-cfp", "90",
		"http://ftp.gwdg.de/pub/linux/misc/packman/suse/openSUSE_Tumbleweed/",
		"packman",
	)
	if err != nil {
		return fmt.Errorf("failed to add Packman repo: %w", err)
	}

	fmt.Println("🔄 Performing vendor switch to Packman for multimedia codecs...")

	err = runCommand(
		"sudo", "zypper", "--non-interactive",
		"dup", "--from", "packman", "--allow-vendor-change",
	)
	if err != nil {
		return fmt.Errorf("failed to perform dup from Packman: %w", err)
	}

	fmt.Println("✅ Codecs installation complete. Multimedia support should be available.")
	return nil
}

func initSudoSession() error {
	fmt.Println("🔐 Asking for sudo password to initialize session...")
	exec.Command("sudo", "-K").Run()
	cmd := exec.Command("sudo", "-v")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to authenticate with sudo: %w", err)
	}

	return nil
}
