package main

import "hyprzen/internal"

func main() {
	internal.PreInstallSetup()
	internal.InstallSystem()
	internal.InstallAps()
	internal.AddConfigs()

}
