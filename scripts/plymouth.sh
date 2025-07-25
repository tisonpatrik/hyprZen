#!/usr/bin/env bash
set -euo pipefail

# === 1. Update GRUB if needed ===
if [ -f /etc/default/grub ]; then
  echo "🔧 Detected GRUB"
  timestamp=$(date +"%Y%m%d%H%M%S")
  sudo cp /etc/default/grub "/etc/default/grub.bak.${timestamp}"

  current_cmdline=$(grep "^GRUB_CMDLINE_LINUX_DEFAULT=" /etc/default/grub | cut -d'"' -f2)
  new_cmdline="$current_cmdline"

  [[ ! "$current_cmdline" =~ splash ]] && new_cmdline+=" splash"
  [[ ! "$current_cmdline" =~ quiet ]] && new_cmdline+=" quiet"
  new_cmdline=$(echo "$new_cmdline" | xargs)

  sudo sed -i "s/^GRUB_CMDLINE_LINUX_DEFAULT=\".*\"/GRUB_CMDLINE_LINUX_DEFAULT=\"$new_cmdline\"/" /etc/default/grub
  echo "🔧 GRUB kernel params updated: $new_cmdline"

  sudo grub2-mkconfig -o /boot/grub2/grub.cfg || sudo grub2-mkconfig -o /boot/efi/EFI/opensuse/grub.cfg
fi

# === 2. Ensure plymouth waits for graphical.target ===
sudo mkdir -p /etc/systemd/system/plymouth-quit.service.d
echo -e "[Unit]\nAfter=graphical.target" | sudo tee /etc/systemd/system/plymouth-quit.service.d/wait-for-graphical.conf

# === 3. Mask plymouth-quit-wait ===
if systemctl list-unit-files | grep -q plymouth-quit-wait.service; then
  sudo systemctl mask plymouth-quit-wait.service
fi

sudo systemctl daemon-reexec
sudo systemctl daemon-reload

# === 4. Install and apply theme ===
if [ -d "$HOME/.local/share/hyprzen/default/plymouth" ]; then
  sudo mkdir -p /usr/share/plymouth/themes/hyprzen
  sudo cp -r "$HOME/.local/share/hyprzen/default/plymouth"/* /usr/share/plymouth/themes/hyprzen/
fi

sudo plymouth-set-default-theme -R hyprzen
echo "✅ plymouth configured with hyprzen theme"
