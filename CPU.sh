#!/usr/bin/env bash
#
# detect-cpu-vendor.sh
# Vrátí intel / amd / other podle výrobce CPU

vendor_id=$(grep -m1 "vendor_id" /proc/cpuinfo | awk '{print $3}')

case "$vendor_id" in
    GenuineIntel)
        echo "intel"
        ;;
    AuthenticAMD)
        echo "amd"
        ;;
    *)
        echo "other"
        ;;
esac

