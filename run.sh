#!/bin/bash
set -e

if command -v go &> /dev/null; then
  echo "Go is already installed: $(go version)"
  exit 0
fi

echo "Installing go:"
sudo zypper in git go

make run
