package cmd

import (
	"fmt"
	"github.com/ddddddO/gtree"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

// ANSI color code for blue text
const blue = "\033[34m"
const reset = "\033[0m"

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all notes in a tree",
	Aliases: []string{"ls"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the root of the tree
		root := gtree.NewRoot("Notes")

		// Map to track the nodes we've already created
		nodeMap := make(map[string]*gtree.Node)
		nodeMap[""] = root // Root node for the top-level directory

		// Walk through the NOTES_PATH directory and build the tree
		err := filepath.Walk(notesPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Ignore hidden files and directories
			components := strings.Split(path, string(os.PathSeparator))
			for _, component := range components {
				if strings.HasPrefix(component, ".") {
					return nil // Skip this file/directory if it's hidden
				}
			}

			// Skip the root itself (".")
			if path == notesPath {
				return nil
			}

			// Get the relative path
			relPath, err := filepath.Rel(notesPath, path)
			if err != nil {
				return err
			}

			// Split the relative path into components (for building the tree structure)
			relComponents := strings.Split(relPath, string(os.PathSeparator))

			// Construct the full path for the parent node (used as a key in the map)
			parentPath := strings.Join(relComponents[:len(relComponents)-1], string(os.PathSeparator))

			// Get the parent node from the map, or default to the root
			parentNode, exists := nodeMap[parentPath]
			if !exists {
				parentNode = root
			}

			// Set folder names to blue
			nodeName := relComponents[len(relComponents)-1]
			if info.IsDir() {
				nodeName = blue + nodeName + reset
			}

			// Add the current component as a node under the parent
			node := parentNode.Add(nodeName)

			// Add the node to the map so that its children can be added later
			nodeMap[relPath] = node

			return nil
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error building tree:", err)
			os.Exit(1)
		}

		// Output the tree
		if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
