#!/bin/bash
# This script, if executed without arguments, will launch your editor and load up your daily note.
# If executed with arguments, the script will append those arguments to your daily note, like this:
# note Eat lunch today.

# Get today's date
today=$(date +"%F")

# Define the path to the .notes_path file in the user's home directory
notes_path="$HOME/.notes_path"

# Check if the .notes_path file exists
if [ ! -e "$notes_path" ]; then
  echo "Error: The .notes_path file does not exist in your home directory."
  exit 1
fi

# Read the path variable from the .notes_path file
dir=$(cat "$notes_path")

# Check if the variable is empty
if [ -z "$dir" ]; then
  echo "Error: The variable in the .notes_path file is empty."
  exit 1
fi

# Create the file name
file_name="$dir/dailies/$today.md"

# Check if the file exists, if not, create it
if [ ! -f "$file_name" ]; then
    touch "$file_name"
fi

# Check if any arguments are given
if [ $# -eq 0 ]; then
    # Open the file with the default editor
    "${EDITOR:-vim}" "$file_name"
else
    # Write each argument to the file on the same line
    echo "- $*" >> "$file_name"
fi
