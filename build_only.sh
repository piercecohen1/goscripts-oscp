#!/bin/bash

# Loop through all subdirectories
for dir in */; do
  if [ -f "$dir/go.mod" ]; then
    # Enter the subdirectory and run go build
    cd "$dir"
    go_files=$(find . -maxdepth 1 -name "*.go")
    for go_file in $go_files; do
      echo "Building ${dir}${go_file#./}"
      go build "$go_file" || { echo "Build failed in ${dir}${go_file#./}"; exit 1; }
    done
    # Return to the parent directory
    cd ..
  fi
done