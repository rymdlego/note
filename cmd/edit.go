package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:     "edit <name>",
	Short:   "Create a new note or edit an existing note",
	Aliases: []string{"insert"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		noteName := args[0] + "." + notesFileType // Ensure the note has a file extension

		// Save the current working directory
		origDir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting current directory: %v", err)
		}

		// Change to NOTES_PATH directory
		if err := os.Chdir(notesPath); err != nil {
			log.Fatalf("Error changing to NOTES_PATH directory: %v", err)
		}

		fmt.Printf("Opening %s in editor: %s\n", noteName, notesEditor)
		editorCmd := exec.Command(notesEditor, noteName) // Launch the editor in the notesPath with the new file
		editorCmd.Stdout = os.Stdout
		editorCmd.Stderr = os.Stderr
		editorCmd.Stdin = os.Stdin

		// Execute the editor and wait for it to finish
		if err := editorCmd.Run(); err != nil {
			log.Fatalf("Error launching editor: %v", err)
		}

		// Restore the original working directory
		if err := os.Chdir(origDir); err != nil {
			log.Fatalf("Error restoring original directory: %v", err)
		}

		fmt.Println("Returned to original directory")
	},
	ValidArgsFunction: CompleteNoteFiles, // Register the file completion function

}

func init() {
	RootCmd.AddCommand(editCmd)
}
