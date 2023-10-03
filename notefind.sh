#!/bin/bash

# Set this to your preferred application for viewing results.
viewer=bat


# Check if there are at least two arguments (keywords and directory)
if [ $# -lt 1 ]; then
  echo "Usage: $0 <keyword1> <keyword2> ..."
  echo "Example: $0 important 2023 # should return notes that contain both 'important' and '2023' in either the note name or its contents."
  exit 1
fi

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

# Check if the provided directory exists
if [ ! -d "$dir" ]; then
  echo "Directory '$dir' does not exist."
  exit 1
fi

# Loop through all text files in the directory and its subdirectories
find "$dir" -type f -name "*.md" | while read -r file; do
  # Initialize a variable to track the number of keyword matches
  match_count=0

  # Loop through the provided keywords and search for them in the file
  for keyword in "$@"; do
    if grep -i -q "$keyword" "$file"; then
      ((match_count++))
    elif [[ "$file" == *"$keyword"* ]]; then
      ((match_count++))
    else
      # If any keyword is not found, break the loop early
      break
    fi
  done

  # If all keywords are found in the file, display the entire file
  if [ "$match_count" -eq "$#" ]; then
    $viewer "$file" # Display the entire file
  fi
done

