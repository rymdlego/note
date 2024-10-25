package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate shell completion script",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			err := RootCmd.GenBashCompletion(os.Stdout)
			if err != nil {
				fmt.Printf("Error generating bash completion script: %v\n", err)
			}
		case "zsh":
			err := RootCmd.GenZshCompletion(os.Stdout)
			if err != nil {
				fmt.Printf("Error generating zsh completion script: %v\n", err)
			}
		case "fish":
			err := RootCmd.GenFishCompletion(os.Stdout, true)
			if err != nil {
				fmt.Printf("Error generating fish completion script: %v\n", err)
			}
		case "powershell":
			err := RootCmd.GenPowerShellCompletion(os.Stdout)
			if err != nil {
				fmt.Printf("Error generating PowerShell completion script: %v\n", err)
			}
		default:
			fmt.Println("Unsupported shell. Please specify bash, zsh, fish, or powershell.")
		}
	},
}

func CompleteNoteFiles(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// Ensure notesPath is set
	if notesPath == "" {
		return nil, cobra.ShellCompDirectiveError
	}

	// Check if there are already arguments provided
	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	var completions []string
	toCompleteLower := strings.ToLower(toComplete)

	// Walk through the notesPath directory
	err := filepath.Walk(notesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip hidden files and folders
		if strings.HasPrefix(info.Name(), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		relPath, err := filepath.Rel(notesPath, path)
		if err != nil {
			return err
		}
		relPathLower := strings.ToLower(relPath)
		if strings.HasPrefix(relPathLower, toCompleteLower) {
			if info.IsDir() {
				completions = append(completions, relPath+"/")
			} else if strings.HasSuffix(info.Name(), "."+notesFileType) {
				completions = append(completions, strings.TrimSuffix(relPath, "."+notesFileType))
			}
		}
		return nil
	})
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	// Return the list of completions
	return completions, cobra.ShellCompDirectiveNoFileComp
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
