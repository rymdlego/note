package cmd

import (
	"fmt"
	// "io/ioutil"
	// "io"
	// "log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find <searchstring>",
	Short: "Find notes that contain the search string (inside files)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		searchString := args[0]
		err := searchNotesInFiles(searchString)
		if err != nil {
			fmt.Printf("Error searching for notes: %v\n", err)
			os.Exit(1)
		}
	},
}

// searchNotesInFiles searches the contents of .md files for the search string
func searchNotesInFiles(searchString string) error {
	// Walk through the NOTES_PATH directory
	err := filepath.Walk(notesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only consider files with .md extension
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			// Read the file contents
			contents, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Search for the string inside the file contents (case-insensitive)
			if strings.Contains(strings.ToLower(string(contents)), strings.ToLower(searchString)) {
				relativePath, _ := filepath.Rel(notesPath, path)
				fmt.Printf("Found in: %s\n", relativePath)
			}
		}
		return nil
	})

	return err
}

func init() {
	RootCmd.AddCommand(findCmd) // Add the "note find" command
}
