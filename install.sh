#!/usr/bin/env sh
# Installs the latest GlassFM release for Linux.
# Usage: curl -fsSL https://raw.githubusercontent.com/Nickaphy/GlassFM/main/install.sh | sh
set -eu

REPO="Nickaphy/GlassFM"
BIN_NAME="glassfm"
INSTALL_DIR="${GLASSFM_INSTALL_DIR:-$HOME/.local/bin}"

if [ "$(uname -s)" != "Linux" ]; then
  echo "This installer only supports Linux." >&2
  exit 1
fi

arch="$(uname -m)"
case "$arch" in
  x86_64|amd64) arch="amd64" ;;
  aarch64|arm64) arch="arm64" ;;
  *)
    echo "Unsupported architecture: $arch" >&2
    exit 1
    ;;
esac

echo "Fetching latest GlassFM release..."
latest_tag="$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" | grep -m1 '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')"

if [ -z "$latest_tag" ]; then
  echo "Could not find a published release for ${REPO}." >&2
  exit 1
fi

archive="glassfm_linux_${arch}.tar.gz"
url="https://github.com/${REPO}/releases/download/${latest_tag}/${archive}"

tmpdir="$(mktemp -d)"
trap 'rm -rf "$tmpdir"' EXIT

echo "Downloading ${archive} (${latest_tag})..."
curl -fsSL "$url" -o "$tmpdir/$archive"

tar -xzf "$tmpdir/$archive" -C "$tmpdir"

mkdir -p "$INSTALL_DIR"
mv "$tmpdir/$BIN_NAME" "$INSTALL_DIR/$BIN_NAME"
chmod +x "$INSTALL_DIR/$BIN_NAME"

echo "Installed glassfm to $INSTALL_DIR/$BIN_NAME"

case ":$PATH:" in
  *":$INSTALL_DIR:"*)
    echo "Run it with: glassfm"
    ;;
  *)
    echo ""
    echo "NOTE: $INSTALL_DIR is not on your PATH."
    echo "Add this to your ~/.bashrc or ~/.zshrc, then restart your shell:"
    echo "  export PATH=\"\$PATH:$INSTALL_DIR\""
    echo "Then run: glassfm"
    ;;
esac
