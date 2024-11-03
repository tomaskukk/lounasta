#!/bin/bash

REPO="tomaskukk/lounasta"
VERSION="0.0.1"  # Replace with the latest release version

# Detect the platform
case "$(uname -s)" in
    Linux*)     OS="linux";;
    Darwin*)    OS="darwin";;
    CYGWIN*|MINGW*|MSYS_NT*) OS="windows";;
    *)          echo "Unsupported OS"; exit 1;;
esac

ARCH="$(uname -m)"
case "$ARCH" in
    x86_64) ARCH="amd64";;
    arm64) ARCH="arm64";;
    *) echo "Unsupported architecture"; exit 1;;
esac

# Download the binary
FILE="lounasta-$OS-$ARCH"
URL="https://github.com/$REPO/releases/download/$VERSION/$FILE.zip"
curl -L -o "$FILE.zip" "$URL" || { echo "Download failed"; exit 1; }

# Extract and install
unzip "$FILE.zip"
chmod +x "$FILE"
mv "$FILE" /usr/local/bin/lounasta || { echo "Permission denied"; exit 1; }

echo "Lounasta installed successfully! Run it with 'lounasta'."
