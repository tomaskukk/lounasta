# Lounasta

**Lounasta** is a cross-platform command-line application that fetches lunch menus from nearby restaurants via the [lounaat.info](https://lounaat.info) API. The app uses the native location API in macOS (Darwin) to determine location and falls back to IP-based location for other platforms.

The **Lounasta** app provides two parameters for filtering lunch menu searches to help refine results based on specific criteria:

1. **Restaurant Name Filter**:

   - **Flag**: `--name` or `-n`
   - **Description**: Filters lunch menus to only include those from restaurants matching the specified name.

2. **Food Name Filter**:
   - **Flag**: `--food` or `-f`
   - **Description**: Filters lunch menus by the type of food. This is useful if you’re looking for a specific dish or ingredient in the menu offerings.

## Features

- Fetches and displays lunch menus from nearby restaurants.
- Supports native location detection on macOS for accurate geolocation.
- Fallback to IP-based location for non-macOS platforms.
- Cross-platform compatibility (macOS, and Linux).

## Installation

Replace version tag with version you want. For 0.0.1 it would be:

```bash
curl -fsSL https://raw.githubusercontent.com/tomaskukk/lounasta/0.0.1/install.sh | bash
```

**Note**
If you use darwin, you'll need to sign the binary first.

This can be done the following way:

Create a signing certificate locally. Instructions can be found [here](https://support.apple.com/guide/keychain-access/create-self-signed-certificates-kyca8916/mac)

Identity type: self signed root.
Certificate type: codesigning.

```bash
  sudo codesign -s "<certificateName>" "$(which lounasta)"
```

## Build Instructions

The build script (`build.sh`) handles cross-compiling for multiple platforms, including macOS (both Intel and ARM architectures), Windows, and Linux. Additionally, the script includes options to:

- Sign the binary for macOS using a specified certificate (Darwin only).
- Install the binary to your Go path for easy access.

### Build Script Usage

Run the script with optional flags:

```bash
./build.sh -c <CERT_NAME> -v -i <INSTALL_OPTION>
```

**Flags:**

- `-c <CERT_NAME>`: Specify the certificate name for signing binaries on macOS.
- `-v`: Enable verbose output during build and signing.
- `-i <INSTALL_OPTION>`: Specify `darwin` or `other` to install the binary to your Go path.

### Build Process Steps

1. **Set Platform Variables:** Compiles for `darwin/amd64`, `darwin/arm64`, `linux/amd64`, and `linux/arm64`.
2. **Build macOS Artifacts:** If targeting macOS, the script compiles the Objective-C files needed for the native location API.
3. **Clean-Up:** Removes intermediate build artifacts.
4. **Signing (macOS only):** If a certificate is provided, the binary will be signed for macOS.

## Dependencies

- **Go:** Ensure Go is installed on your machine.
- **Xcode (macOS):** Required to build macOS location artifacts.

## Installation

To install and access `lounasta` globally, use the `-i` flag with `darwin` or `other`, depending on your platform.
