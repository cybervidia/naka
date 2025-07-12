#!/bin/bash

set -e

APP_NAME="naka"
SRC_FILE="main.go"
RELEASE_DIR="Releases"
FILES_TO_INCLUDE="README.md LICENSE"

PLATFORMS=(
  "linux amd64 tar.gz"
  "windows amd64 zip"
  "darwin amd64 tar.gz"
  "darwin arm64 tar.gz"
)

# Cleanup previous builds
rm -rf "$RELEASE_DIR"
mkdir -p "$RELEASE_DIR"

build() {
  local GOOS="$1"
  local GOARCH="$2"
  local ARCHIVE_TYPE="$3"
  local BUILD_DIR="$RELEASE_DIR/${GOOS}_${GOARCH}"
  local OUTPUT_NAME="$APP_NAME"

  [ "$GOOS" == "windows" ] && OUTPUT_NAME="$APP_NAME.exe"

  echo "▶️ Building for $GOOS/$GOARCH..."
  mkdir -p "$BUILD_DIR"
  env GOOS=$GOOS GOARCH=$GOARCH go build -o "$BUILD_DIR/$OUTPUT_NAME" "$SRC_FILE"

  # Include extra files
  for f in $FILES_TO_INCLUDE; do
    [ -f "$f" ] && cp "$f" "$BUILD_DIR/"
  done

  # Archive
  pushd "$RELEASE_DIR" > /dev/null
  local ARCHIVE_NAME="${APP_NAME}-${GOOS}-${GOARCH}.${ARCHIVE_TYPE}"
  if [ "$ARCHIVE_TYPE" == "zip" ]; then
    zip -r "$ARCHIVE_NAME" "${GOOS}_${GOARCH}" > /dev/null
  else
    tar -czf "$ARCHIVE_NAME" "${GOOS}_${GOARCH}"
  fi
  popd > /dev/null

  echo "✅ $ARCHIVE_NAME created."
}

# Build for each target
for entry in "${PLATFORMS[@]}"; do
  read -r GOOS GOARCH ARCHIVE_TYPE <<< "$entry"
  build "$GOOS" "$GOARCH" "$ARCHIVE_TYPE"
done

# Create macOS universal binary
if [[ "$OSTYPE" == "darwin"* ]]; then
  echo "🛠 Creating macOS universal binary with lipo..."

  DARWIN_DIR="$RELEASE_DIR/darwin_universal"
  mkdir -p "$DARWIN_DIR"

  lipo -create \
    "$RELEASE_DIR/darwin_amd64/$APP_NAME" \
    "$RELEASE_DIR/darwin_arm64/$APP_NAME" \
    -output "$DARWIN_DIR/$APP_NAME"

  # Copy README and LICENSE
  for f in $FILES_TO_INCLUDE; do
    [ -f "$f" ] && cp "$f" "$DARWIN_DIR/"
  done

  pushd "$RELEASE_DIR" > /dev/null
  tar -czf "${APP_NAME}-darwin-universal.tar.gz" "darwin_universal"
  popd > /dev/null

  echo "✅ ${APP_NAME}-darwin-universal.tar.gz created."
else
  echo "⚠️ Skipping macOS universal build: not running on macOS (lipo not available)."
fi

echo "🎉 All builds complete. Check the '$RELEASE_DIR' folder."
