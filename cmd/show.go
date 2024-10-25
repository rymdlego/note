package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:     "show",
	Short:   "Display your note",
	Aliases: []string{"view", "display"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		noteName := args[0] + "." + notesFileType // Ensure the note has a file extension

		fmt.Printf("Opening %s in viewer: %s\n", noteName, notesViewer)
		viewerCmd := exec.Command(notesViewer, notesPath+"/"+noteName) // Launch the editor in the notesPath with the new file
		viewerCmd.Stdout = os.Stdout
		viewerCmd.Stderr = os.Stderr
		viewerCmd.Stdin = os.Stdin

		// Execute the viewer and wait for it to finish
		if err := viewerCmd.Run(); err != nil {
			log.Fatalf("Error launching viewer: %v", err)
		}
	},
	ValidArgsFunction: CompleteNoteFiles, // Register the file completion function
}

func init() {
	RootCmd.AddCommand(showCmd)
}
