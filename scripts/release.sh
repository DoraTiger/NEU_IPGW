#!/usr/bin/env bash

set -euo pipefail

name=$1
targets_dir="$PWD/$2"
release_dir="$PWD/$3"

# SHA256 function
if command -v sha256sum >/dev/null 2>&1; then
  sha256() {
    sha256sum "$1"
  }
elif command -v shasum >/dev/null 2>&1; then
  sha256() {
    shasum -a 256 "$1"
  }
else
  echo "ERROR: No SHA256 tool found. Please install 'coreutils' (Linux) or use macOS with 'shasum'." >&2
  exit 1
fi

# check args
if [ -z "$name" ] || [ -z "$targets_dir" ] || [ -z "$release_dir" ]; then
  echo "Usage: $0 <name> <targets_dir> <release_dir>"
  exit 1
fi

echo "start: zip all targets in $targets_dir to $release_dir."
mkdir -p "$release_dir"

checksums_temp=$(mktemp)

while IFS= read -r -d $'\0' dir; do
  target_name=$(basename "$dir")
  zip_file="$release_dir/${name}-${target_name}.tar.gz"

  echo "zip $target_name to $zip_file"
  tar -czf "$zip_file" -C "$dir" .

  echo "Generated: $zip_file.sha256"
  sha256 "$zip_file" | sed "s| .*/| |" > "$zip_file.sha256"

  cat "$zip_file.sha256" >> "$checksums_temp"

done < <(find "$targets_dir" -mindepth 1 -maxdepth 1 -type d -print0)

if [ -s "$checksums_temp" ]; then
  sort "$checksums_temp" > "$release_dir/checksums.txt"
  echo "Generated: $release_dir/checksums.txt"
else
  echo "WARNING: No platform directories found in $targets_dir"
fi

rm -f "$checksums_temp"

echo "done."
