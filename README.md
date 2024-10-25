# Note

**Note** is a command-line interface tool for managing notes efficiently. It supports creating, editing, viewing, searching notes, and includes Git integration to keep your notes synchronized and version-controlled. It works amazingly well with note systems like Obsidian.

## Features

- Create and edit notes using your preferred text editor.
- Automatically save notes in the specified notes directory (`NOTES_PATH`).
- Supports searching notes by name and content.
- Git integration for version control
- Scratch notes functionality for quick note-taking.
- Shell autocompletion for commands and file names (supports Bash, Zsh, Fish, and PowerShell).

## Quicky install (requires Go):

```bash
go install github.com/rymdlego/go@latest
```

This will place the executable in $HOME/go/bin or $GOPATH (if you have it set to something else).
Make sure to have that in your $PATH

## The Git clone install (requires Go)

1. Clone the repository:

   ```bash
   git clone https://github.com/rymdlego/note.git
   cd note
   ```

2. Build the application:

   ```bash
   go build -o note
   ```

3. Move the `note` binary to your PATH:
   ```bash
   sudo mv note /usr/local/bin/
   ```

## Setup

Before using the `note` CLI, ensure that the following environment variables are set:

- `NOTES_PATH`: Path to the directory where notes are stored.
- `EDITOR`: The text editor to use when editing notes (default: `vi`).
- `NOTES_FILETYPE`: The file type for notes (default: `md`).
- `NOTES_SCRATCHFILE`: The scratch file name for quick notes (default: `scratch`).

For example:

```bash
export NOTES_PATH="$HOME/notes"
export EDITOR="nvim"
```

## Usage

### General Commands

1. **Edit or Create a Note**:

   ```bash
   note edit <name>
   ```

   This command opens the note for editing. If the note doesn’t exist, it creates a new one with the given name.

   Alias: `note insert <name>`

2. **View a Note**:

   ```bash
   note show <name>
   ```

   Displays the contents of the note using the configured viewer (`NOTES_VIEWER`, default is `cat`).

   Aliases: `note view`, `note display`

3. **List Notes**:

   ```bash
   note list
   ```

   List all notes in a tree structure.

   Alias: `note ls`

4. **Search Notes by Name or Content**:

   ```bash
   note find <searchstring>
   ```

   Finds notes whose names or content contain the search string.

5. **Version**:

   ```bash
   note version
   ```

   Displays the current version of the application.

6. **Git Integration**:
   You can run any Git command inside your notes directory.

   ```bash
   note git <git-command>
   ```

   Examples:

   - `note git status`
   - `note git add .`
   - `note git commit -m "Updated notes"`

   You can also use shortcuts for common Git commands:

   - `note status`
   - `note add <files>`
   - `note commit -m "message"`
   - `note push`
   - `note pull`

### Scratch Notes

If you run `note` without any arguments, it opens a "scratch" note for quick notes:

```bash
note
```

You can also pipe content to the scratch note:

```bash
echo "Eggs and Avocadoes" | note
```

## Autocompletion

To enable shell autocompletion, use the following command based on your shell:

- **Bash**:

  ```bash
  source <(note completion bash)
  ```

- **Zsh**:

  ```bash
  source <(note completion zsh)
  ```

- **Fish**:

  ```bash
  note completion fish | source
  ```

- **PowerShell**:
  ```bash
  note completion powershell | Out-String | Invoke-Expression
  ```

To permanently enable autocompletion, add the corresponding command to your shell’s configuration file (`~/.bashrc`, `~/.zshrc`, etc.).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
