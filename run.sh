#!/bin/bash

echo "Starting build process..."

# Define variables
APP_NAME="lounasta"
APP_BUNDLE="$APP_NAME.app"
APP_EXECUTABLE="$APP_BUNDLE/Contents/MacOS/$APP_NAME"
INFO_PLIST="$APP_BUNDLE/Contents/Info.plist"
CERT_NAME="Developer Certificate"  # Use the name from your certificate

# Remove old build artifacts
rm -rf "$APP_BUNDLE" ./location_manager/location_manager_darwin.o liblocation.a

# Build the C library
clang -c -arch arm64 -o location_manager/location_manager_darwin.o location_manager/location_manager_darwin.m
ar rcs liblocation.a location_manager/location_manager_darwin.o

# Create the app bundle structure
mkdir -p "$APP_BUNDLE/Contents/MacOS"

# Copy Info.plist into the bundle
cp location_manager/Info.plist "$INFO_PLIST"

# Build the Go binary
go build -gcflags "all=-N -l" -ldflags="-linkmode=external" -o "$APP_EXECUTABLE" -v -x main.go

# Sign the app bundle
codesign -s "$CERT_NAME" --deep --force --verbose "$APP_BUNDLE"

echo "Build process completed."

# Run the app
open "$APP_BUNDLE"
