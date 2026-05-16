#!/bin/bash
set -e

# Rename tag files to index.ts in services directories.
# orval tags-split mode creates: services/tagName/tagName.ts
# We rename to:                  services/tagName/index.ts
echo "Renaming tag files to index.ts..."
find src/api/generated/services -mindepth 1 -maxdepth 1 -type d | while read -r dir; do
  dirname=$(basename "$dir")
  tagfile="$dir/$dirname.ts"
  if [ -f "$tagfile" ]; then
    mv "$tagfile" "$dir/index.ts"
    echo "  $tagfile -> $dir/index.ts"
  fi
done
echo "Done."

# Format generated files with prettier.
echo "Formatting generated files..."
npx prettier --write src/api/generated
echo "Done."

echo "API client generation complete."
