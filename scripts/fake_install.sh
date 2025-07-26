#!/bin/bash
set -euo pipefail

echo "🚀 Starting fake installation process..."

# Simulate some work by waiting 3 seconds
echo "⏳ Simulating installation work..."
sleep 3

echo "🔧 Requesting sudo privileges for fake system modification..."

# This will trigger a real sudo prompt
# We'll use a harmless command that requires sudo but doesn't actually change anything
sudo true

# Simulate some more work after sudo
echo "✅ Sudo authentication successful"
echo "🔧 Performing fake system modifications..."
sleep 1

# Create a temporary file to simulate some work
sudo touch /tmp/fake_install_test_$(date +%s)

echo "🎉 Fake installation completed successfully!"
echo "�� Created test file: /tmp/fake_install_test_*"

# Clean up the test file
sudo rm -f /tmp/fake_install_test_*

echo "🧹 Cleanup completed"
echo "✅ Fake installation process finished successfully"

exit 0 