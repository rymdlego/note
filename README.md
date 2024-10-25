# Note

**Note** is a simple but amazing note-taking tool for the command, written in Go. It supports creating, editing, viewing, searching notes, and includes Git integration to keep your notes synchronized and version-controlled. It works amazingly well with note systems like Obsidian.

## Features

- Create and edit notes using your preferred text editor.
- Scratch note functionality for instant note-taking, and piping output into a note. (`echo !! | note`)
- Automatically save notes in the specified notes directory (`NOTES_PATH`).
- Supports searching notes by their content to easily locate what you need.
- Git integration for version control
- Shell autocompletion for commands AND file names (supports Bash, Zsh, Fish, and PowerShell).

## Quicky install (requires Go):

```bash
go install github.com/rymdlego/note@latest
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

Before using the `note` CLI, ensure that the following environment variable is set:

- `NOTES_PATH`: Path to the directory where notes are stored.

Optional environment variables:

- `EDITOR`: The text editor to use when editing notes (default: `vi`).
- `NOTES_FILETYPE`: The file type for notes (default: `md`).
- `NOTES_SCRATCHFILE`: The scratch file name for quick notes (default: `scratch`).

For example:

```bash
export NOTES_PATH="$HOME/notes"
export EDITOR="nvim"
export NOTES_SCRATCHFILE="random"
```

## Usage

### General Commands

1. **Edit or Open a new Note**:

   ```bash
   note edit <name>
   ```

   This command opens the note for editing, or just opens the editor with the desired document name.

   Alias: `note insert <name>`

2. **View a Note**:

   ```bash
   note show <name>
   ```

   Displays the contents of the note using the configured viewer (`NOTES_VIEWER`, default is `cat`, you might want to use something cool like `glow`).

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

   Finds notes whose content contain the search string.

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

   You can also use these shorthand commands for common Git actions:

   - `note status`
   - `note add <files>`
   - `note commit -m "message"`
   - `note push`
   - `note pull`

### Scratch Note

If you run `note` without any arguments, it opens the "scratch" note, which is just a generic note for unsorted ideas:

```bash
note
```

You can also pipe content to the scratch note:

```bash
echo "Eggs and Avocadoes" | note
```

Why not document that long command you just ran:

```bash
kubectl create secret generic something -n default --from-env-file myfile --dry-run=client -o yaml
echo !! | note
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
