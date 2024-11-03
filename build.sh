#!/bin/bash

echo "Starting build process..."

# Define variables
PACKAGE_NAME="github.com/tomaskukk/lounasta"
APP_NAME="lounasta"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
LOCATION_PROVIDER_DARWIN_DIR="$SCRIPT_DIR/location/location_provider_darwin"
INFO_PLIST="$LOCATION_PROVIDER_DARWIN_DIR/Info.plist"
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


clean_up_c_libraries() {
  rm -rf $VERBOSE $SCRIPT_DIR/$APP_NAME $LOCATION_PROVIDER_DARWIN_DIR/location_provider_darwin.o $LOCATION_PROVIDER_DARWIN_DIR/liblocation.a
}

clean_up_build_dir() {
  rm -rf $VERBOSE $SCRIPT_DIR/build
}

clean_up_build_dir


platforms=("windows/amd64" "darwin/amd64" "darwin/arm64" "linux/amd64" "linux/arm64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	OUTPUT_NAME='build/'$APP_NAME'-'$GOOS'-'$GOARCH
  CGO_LDFLAGS=''
	if [ $GOOS = "windows" ]; then
		OUTPUT_NAME+='.exe'
	fi

  # Set up Darwin-specific flags
  if [ "$GOOS" = "darwin" ]; then
    if [ "$GOARCH" = "arm64" ]; then
      C_ARCH="arm64"
    else
      C_ARCH="x86_64"
    fi

    clang -c -arch $C_ARCH -o "$LOCATION_PROVIDER_DARWIN_DIR/location_provider_darwin.o" "$LOCATION_PROVIDER_DARWIN_DIR/location_provider_darwin.m"
    ar rcs "$LOCATION_PROVIDER_DARWIN_DIR/liblocation.a" "$LOCATION_PROVIDER_DARWIN_DIR/location_provider_darwin.o"

    CGO_ENABLED=1
    LDFLAGS="-extldflags \"-L$LOCATION_PROVIDER_DARWIN_DIR -sectcreate __TEXT __info_plist $INFO_PLIST\" -linkmode=external"
  else
    LDFLAGS=""
    CGO_ENABLED=0
  fi

  echo $LDFLAGS

  env CGO_ENABLED=$CGO_ENABLED GOOS=$GOOS GOARCH=$GOARCH go build -gcflags "all=-N -l" -ldflags="$LDFLAGS" -o "$OUTPUT_NAME" $VERBOSE -x "$PACKAGE_NAME"  

  clean_up_c_libraries

  if [ $GOOS = "darwin" ]; then
    if [ -n "$CERT_NAME" ]; then
      codesign $VERBOSE -s "$CERT_NAME" "./$OUTPUT_NAME"
    else
      codesign $VERBOSE -s - "./$OUTPUT_NAME"
    fi
  fi

	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done

echo "Build process completed."
