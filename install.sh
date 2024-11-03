#!/bin/bash

REPO="tomaskukk/lounasta"
VERSION="0.0.1"  # Replace with the latest release version

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
    x86_64) ARCH="amd64";;
    arm64) ARCH="arm64";;
    *) echo "Unsupported architecture: $ARCH"; exit 1;;
esac


FILE="lounasta-$OS-$ARCH"
URL="https://github.com/$REPO/releases/download/$VERSION/$FILE.zip"

# Download and install
echo "Downloading $FILE from $URL..."
curl -L -o "$FILE.zip" "$URL" || { echo "Download failed"; exit 1; }

echo "Extracting $FILE.zip..."
unzip -o "$FILE.zip" || { echo "Extraction failed"; exit 1; }
chmod +x "$FILE"

echo "Installing $FILE to /usr/local/bin..."
sudo mv "$FILE" /usr/local/bin/lounasta || { echo "Installation failed"; exit 1; }

echo "Cleaning up..."
rm "$FILE.zip"

echo "Lounasta installed successfully! You can now run it with 'lounasta'."
