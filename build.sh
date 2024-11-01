#!/bin/bash

# Your build script commands go here

echo "Starting build process..."

# Remove old build artifacts
rm -f ./lounasta ./location_manager_darwin.o liblocation.a

# Build the C library
clang -c -arch arm64 -o location_manager_darwin.o location_manager_darwin.m
ar rcs liblocation.a location_manager_darwin.o

# Build the Go binary
go build -gcflags "all=-N -l" -ldflags="-extldflags \"-sectcreate __TEXT __info_plist $(pwd)/Info.plist\" -linkmode=external" -o lounasta -v -x main.go

# Sign the binary
codesign -s - lounasta

# Run the binary
echo "Build process completed."

./lounasta "$@"
