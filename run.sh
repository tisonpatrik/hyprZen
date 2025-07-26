#!/bin/bash
set -e

sudo pacman -Sy --noconfirm --needed go

clear

make build

./build/main