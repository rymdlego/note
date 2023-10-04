# note
Access your Obsidian notes from the terminal
## Installation
Copy the `note` script into a pathed folder, for instance `~/bin`
Set an environment variable for `NOTES_PATH`, like this (no trailing slash):
```
export NOTES_PATH=/path/to/my/obsidian/vault
```
### There are some more optional things to set:
```
export NOTES_DAILIES=dailies
```
This sets a subfolder for daily notes (no trailing slash). In this example they will be placed in /dailies/

```
export NOTES_PREFIX='- '
```
When you quick-append to daily note, the rows will be prefixed with whatever you specify here (this example turns these rows into dotted list). Default is blank.

```
export NOTES_VIEWER=bat
```
If you want to use something else but cat for displaying notes. I recommend bat. It's very nice for this. Default is cat.

```
export EDITOR=vim
```
Set the default editor in your shell. This is probably already in place.

```
alias n=note
```
Do it.

## Usage
```
# n
```
This launches your default editor and loads up your daily note.

```
# n something very important
```
This appends *something very important* to your daily note.

```
# openssl s_client -connect example.com:443 -servername example.com -showcerts </dev/null | openssl x509 -noout -dates
# n !!
```
Just another example of appending information. You run some complicated command and you want to remember it for future use. `n !!`will add the previous command to your daily note.

```
# n -s vpn customer infra
```
The `-s` flag will search your notes for the keywords that you specify. In this case, it will search for `vpn` `customer` and `infra`. It will display the notes that contain all of these words (it searches both filenames and the file contents).

```
# n -f books
```
The `-f` flag will open up note files in your default editor. In this case we will edit books.md. If the file doesn't exist, it will be created.

```
# n -f
```
This one will work if you have `fzf` installed. By running `n -f` with no other arguments, you will be able to fuzzy find any document in `fzf` fashion. I highly recommend it.

```
# n -v kubernetes
```
The `-v` flag will view notes. In this case we will view kubernetes.md. If you specify more notes, those will be displayed as well.
