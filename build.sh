#!/bin/bash

echo "Starting build process..."

# Define variables
APP_NAME="lounasta"
CURRENT_DIR=$(pwd)
LOCATION_MANAGER_DIR="$CURRENT_DIR/location_manager"
EXECUTABLE_PATH="$CURRENT_DIR/$APP_NAME"
INFO_PLIST="$LOCATION_MANAGER_DIR/Info.plist"
CERT_NAME=""
VERBOSE=""

# Parse arguments
while getopts "c:v" opt; do
  case $opt in
    c)
      CERT_NAME=$OPTARG
      ;;
    v)
      VERBOSE="-v"
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done

# Remove old build artifacts
rm -rf $VERBOSE $CURRENT_DIR/$APP_NAME $CURRENT_DIR/location_manager/location_manager_darwin.o $LOCATION_MANAGER_DIR/liblocation.a

# Build the C library
clang -c -arch arm64 -o $LOCATION_MANAGER_DIR/location_manager_darwin.o $LOCATION_MANAGER_DIR/location_manager_darwin.m

ar rcs $LOCATION_MANAGER_DIR/liblocation.a $LOCATION_MANAGER_DIR/location_manager_darwin.o

# Build the Go binary
go build -gcflags "all=-N -l" -ldflags="-extldflags \"-L$LOCATION_MANAGER_DIR -sectcreate __TEXT __info_plist $INFO_PLIST\" -linkmode=external" -o $APP_NAME $VERBOSE -x "main.go"

# Sign the exectuable
if [ -n "$CERT_NAME" ]; then
  codesign $VERBOSE -s "$CERT_NAME" "$EXECUTABLE_PATH"
else
  codesign $VERBOSE -s - "$EXECUTABLE_PATH"
fi

echo "Build process completed."
