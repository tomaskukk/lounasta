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
INSTALL=""

while getopts "c:vi:" opt; do
  case $opt in
    c)
      echo "Certificate Name: $OPTARG"
      CERT_NAME=$OPTARG
      ;;
    v)
      VERBOSE="-v"
      ;;
    i)
      if [[ "$OPTARG" == "darwin" || "$OPTARG" == "other" ]]; then
        echo "Adding install flag. Binary will be installed to gopath."
        INSTALL=$OPTARG
      else
        echo "Invalid value for -i. Use 'darwin' or 'other'."
        exit 1
      fi
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done

sign_binary_for_darwin() {
  local output_name=$1
  echo "Signing binary $output_name for Darwin..."
  if [ -n "$CERT_NAME" ]; then
    codesign $VERBOSE -s "$CERT_NAME" "$output_name"
  else
    codesign $VERBOSE -s - "$output_name"
  fi
}

build_c_artifacts() {
  clang -c -arch $C_ARCH -o "$LOCATION_PROVIDER_DARWIN_DIR/location_provider_darwin.o" "$LOCATION_PROVIDER_DARWIN_DIR/location_provider_darwin.m"
  ar rcs "$LOCATION_PROVIDER_DARWIN_DIR/liblocation.a" "$LOCATION_PROVIDER_DARWIN_DIR/location_provider_darwin.o"
}

clean_up_c_artifacts() {
  rm -rf $VERBOSE $SCRIPT_DIR/$APP_NAME $LOCATION_PROVIDER_DARWIN_DIR/location_provider_darwin.o $LOCATION_PROVIDER_DARWIN_DIR/liblocation.a
}

clean_up_build_dir() {
  rm -rf $VERBOSE $SCRIPT_DIR/build
}

clean_up_build_dir


platforms=("darwin/amd64" "darwin/arm64" "linux/amd64" "linux/arm64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	OUTPUT_NAME=$SCRIPT_DIR'/build/'$APP_NAME'-'$GOOS'-'$GOARCH
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

    build_c_artifacts

    CGO_ENABLED=1
    LDFLAGS="-extldflags \"-L$LOCATION_PROVIDER_DARWIN_DIR -sectcreate __TEXT __info_plist $INFO_PLIST\" -linkmode=external"
  else
    LDFLAGS=""
    CGO_ENABLED=0
  fi

  echo "Building for $GOOS/$GOARCH..."

  env CGO_ENABLED=$CGO_ENABLED GOOS=$GOOS GOARCH=$GOARCH go build -gcflags "all=-N -l" -ldflags="$LDFLAGS" -o "$OUTPUT_NAME" $VERBOSE "$PACKAGE_NAME"

  clean_up_c_artifacts

  if [ $GOOS = "darwin" ]; then
    sign_binary_for_darwin $OUTPUT_NAME
  fi

	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done

echo "Build process completed."


if [ -n "$INSTALL" ]; then
  echo "Installing the binary locally..."
  if [ "$INSTALL" = "darwin" ]; then
    build_c_artifacts
    LDFLAGS="-extldflags \"-L$LOCATION_PROVIDER_DARWIN_DIR -sectcreate __TEXT __info_plist $INFO_PLIST\" -linkmode=external"
  else
    LDFLAGS=""
  fi

  go install -gcflags "all=-N -l" -ldflags="$LDFLAGS" $VERBOSE "$PACKAGE_NAME"

  if [ $INSTALL = "darwin" ]; then
    sign_binary_for_darwin "$(which $APP_NAME)"
    clean_up_c_artifacts
  fi
fi
