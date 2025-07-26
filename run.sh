#!/bin/bash
set -e

sudo pacman -Sy --noconfirm --needed go

make build

./build/main