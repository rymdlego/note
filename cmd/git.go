package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Use:                "git",
	Short:              "Git operations for your notes",
	Args:               cobra.ArbitraryArgs, // Accept arbitrary arguments, including flags
	DisableFlagParsing: true,                // Disable Cobra flag parsing for this command
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure that some git command was passed
		if len(args) == 0 {
			fmt.Println("Please provide a git command (e.g., status, pull, push)")
			return
		}

		// Build the git command using the passed arguments
		gitCmd := exec.Command("git", args...)

		// Set the working directory to the NOTES_PATH
		gitCmd.Dir = filepath.Join(notesPath)

		// Redirect the output of the git command to the terminal
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		gitCmd.Stdin = os.Stdin

		// Run the git command
		err := gitCmd.Run()
		if err != nil {
			log.Fatalf("Error running git command: %v", err)
		}
	},
}

// statusCmd represents the shortcut for "git status"
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the working tree status",
	Run: func(cmd *cobra.Command, args []string) {
		// Run "git status"
		gitCmd := exec.Command("git", "status")
		gitCmd.Dir = filepath.Join(notesPath)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		gitCmd.Stdin = os.Stdin

		// Execute git status
		err := gitCmd.Run()
		if err != nil {
			log.Fatalf("Error running git status: %v", err)
		}
	},
}

// pullCmd represents the shortcut for "git pull"
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Fetch from and integrate with another repository or a local branch",
	Run: func(cmd *cobra.Command, args []string) {
		// Run "git pull"
		gitCmd := exec.Command("git", "pull")
		gitCmd.Dir = filepath.Join(notesPath)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		gitCmd.Stdin = os.Stdin

		// Execute git pull
		err := gitCmd.Run()
		if err != nil {
			log.Fatalf("Error running git pull: %v", err)
		}
	},
}

// pushCmd represents the shortcut for "git push"
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push your commits to remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		// Run "git push"
		gitCmd := exec.Command("git", "push")
		gitCmd.Dir = filepath.Join(notesPath)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		gitCmd.Stdin = os.Stdin

		// Execute git push
		err := gitCmd.Run()
		if err != nil {
			log.Fatalf("Error running git push: %v", err)
		}
	},
}

// addCmd represents the shortcut for "git add"
var addCmd = &cobra.Command{
	Use:   "add [files...]",
	Short: "Add files to Git staging",
	Args:  cobra.MinimumNArgs(1), // Ensure at least one argument (e.g., files or directories) is provided
	Run: func(cmd *cobra.Command, args []string) {
		// Append the git "add" command with the provided args (e.g., "git add .")
		gitArgs := append([]string{"add"}, args...)
		gitCmd := exec.Command("git", gitArgs...)

		// Set the working directory to NOTES_PATH
		gitCmd.Dir = filepath.Join(notesPath)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		gitCmd.Stdin = os.Stdin

		// Execute git add with arguments
		err := gitCmd.Run()
		if err != nil {
			log.Fatalf("Error running git add: %v", err)
		}
	},
}

// commitCmd represents the shortcut for "git commit"
var commitCmd = &cobra.Command{
	Use:                "commit [flags]",
	Short:              "Commit changes to the Git repository",
	Args:               cobra.ArbitraryArgs, // Accept arbitrary arguments (e.g., -m, files)
	DisableFlagParsing: true,                // Disable Cobra flag parsing for this command
	Run: func(cmd *cobra.Command, args []string) {
		// If no arguments are passed, simply run "git commit"
		gitArgs := append([]string{"commit"}, args...) // Append any additional flags or arguments

		// Build the git commit command with the provided arguments
		gitCmd := exec.Command("git", gitArgs...)

		// Set the working directory to NOTES_PATH
		gitCmd.Dir = filepath.Join(notesPath)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		gitCmd.Stdin = os.Stdin

		// Execute git commit with arguments
		err := gitCmd.Run()
		if err != nil {
			log.Fatalf("Error running git commit: %v", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(gitCmd)
	RootCmd.AddCommand(statusCmd)
	RootCmd.AddCommand(pullCmd)
	RootCmd.AddCommand(pushCmd)
	RootCmd.AddCommand(addCmd)
	RootCmd.AddCommand(commitCmd)
}
