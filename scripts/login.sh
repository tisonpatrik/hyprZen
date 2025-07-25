#!/usr/bin/env bash
set -euo pipefail

# === 1. Compile seamless-login binary ===
cat <<'EOF' >/tmp/seamless-login.c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <linux/kd.h>
#include <linux/vt.h>
#include <string.h>

int main(int argc, char *argv[]) {
    int vt_fd;
    int vt_num = 1;
    char vt_path[32];

    if (argc < 2) {
        fprintf(stderr, "Usage: %s <session_command>\n", argv[0]);
        return 1;
    }

    snprintf(vt_path, sizeof(vt_path), "/dev/tty%d", vt_num);
    vt_fd = open(vt_path, O_RDWR);
    if (vt_fd < 0) {
        perror("Failed to open VT");
        return 1;
    }

    if (ioctl(vt_fd, VT_ACTIVATE, vt_num) < 0) {
        perror("VT_ACTIVATE failed");
        close(vt_fd);
        return 1;
    }

    if (ioctl(vt_fd, VT_WAITACTIVE, vt_num) < 0) {
        perror("VT_WAITACTIVE failed");
        close(vt_fd);
        return 1;
    }

    if (ioctl(vt_fd, KDSETMODE, KD_GRAPHICS) < 0) {
        perror("KDSETMODE KD_GRAPHICS failed");
        close(vt_fd);
        return 1;
    }

    const char *clear_seq = "\33[H\33[2J";
    write(vt_fd, clear_seq, strlen(clear_seq));
    close(vt_fd);

    const char *home = getenv("HOME");
    if (home) chdir(home);

    execvp(argv[1], &argv[1]);
    perror("Failed to exec session");
    return 1;
}
EOF

gcc -o /tmp/seamless-login /tmp/seamless-login.c
sudo mv /tmp/seamless-login /usr/local/bin/seamless-login
sudo chmod +x /usr/local/bin/seamless-login
rm /tmp/seamless-login.c

# === 2. Create systemd service ===
cat <<EOF | sudo tee /etc/systemd/system/hyprzen-seamless-login.service
[Unit]
Description=hyprzen seamless auto-login
Documentation=https://github.com/tisonpatrik/hyprzen
Conflicts=getty@tty1.service
After=systemd-user-sessions.service getty@tty1.service plymouth-quit.service systemd-logind.service
PartOf=graphical.target

[Service]
Type=simple
ExecStart=/usr/local/bin/seamless-login uwsm start -- hyprland.desktop
Restart=always
RestartSec=2
User=$(id -un)
TTYPath=/dev/tty1
TTYReset=yes
TTYVHangup=yes
TTYVTDisallocate=yes
StandardInput=tty
StandardOutput=journal
StandardError=journal+console
PAMName=login

[Install]
WantedBy=graphical.target
EOF

# === 3. Wait for graphical.target before quitting plymouth ===
sudo mkdir -p /etc/systemd/system/plymouth-quit.service.d
cat <<EOF | sudo tee /etc/systemd/system/plymouth-quit.service.d/wait-for-graphical.conf 
[Unit]
After=multi-user.target
EOF

# === 4. Mask plymouth-quit-wait.service ===
sudo systemctl mask plymouth-quit-wait.service

# === 5. Enable login service and disable getty ===
sudo systemctl daemon-reexec
sudo systemctl daemon-reload
sudo systemctl enable hyprzen-seamless-login.service
sudo systemctl disable getty@tty1.service

echo "\u2705 hyprzen seamless-login service created and enabled."
