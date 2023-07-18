#!/bin/bash

# Directory where this script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

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

# Check which shell the user is using and write to the corresponding shell config file
if [ -n "$BASH_VERSION" ]; then
    echo "alias urlencode='${SCRIPT_DIR}/url_encode/url_encode'" >> ~/.bashrc
    echo "alias escape_double_quotes='${SCRIPT_DIR}/escape_quotes/escape_double_quotes'" >> ~/.bashrc
    echo "alias psencode='${SCRIPT_DIR}/base64_powershell/b64'" >> ~/.bashrc
    echo "alias path='${SCRIPT_DIR}/path/copy_path'" >> ~/.bashrc
    echo "Success! Please source your bash configuration file with 'source ~/.bashrc'"
elif [ -n "$ZSH_VERSION" ]; then
    echo "alias urlencode='${SCRIPT_DIR}/url_encode/url_encode'" >> ~/.zshrc
    echo "alias escape_double_quotes='${SCRIPT_DIR}/escape_quotes/escape_double_quotes'" >> ~/.zshrc
    echo "alias psencode='${SCRIPT_DIR}/base64_powershell/b64'" >> ~/.zshrc
    echo "alias path='${SCRIPT_DIR}/path/copy_path'" >> ~/.zshrc
    echo "Success! Please source your zsh configuration file with 'source ~/.zshrc'"
else
    echo "Sorry, we don't support your shell."
fi

