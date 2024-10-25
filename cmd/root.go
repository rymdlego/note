package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	notesScratchFile string
	notesPath        string
	notesFileType    string
	notesEditor      string
	notesViewer      string
)

// RootCmd is the base command
var RootCmd = &cobra.Command{
	Use:   "note",
	Short: "note is a CLI for managing notes",
	Run: func(cmd *cobra.Command, args []string) {
		// Check for piped input
		fileInfo, err := os.Stdin.Stat()
		if err != nil {
			log.Fatalf("Error stating stdin: %v", err)
		}

		// If there is piped input, append to the scratch file
		if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
			appendToScratchFile()
		} else {
			// If no input and no args, open the scratch file for editing
			openScratchFileInEditor()
		}
	},
}

// Execute starts the root command and runs the CLI
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Helper function to get an environment variable with a default value
func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Helper function to check if a directory exists, and create it if necessary
func ensureDirectoryExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("NOTES_PATH directory %s does not exist. Creating it...\n", path)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatalf("Error creating NOTES_PATH directory: %v", err)
		}
	}
}

// Ensure NOTES_FILETYPE, NOTES_PATH, and NOTES_SCRATCHFILE are set
func init() {
	notesFileType = getEnvWithDefault("NOTES_FILETYPE", "md")
	notesPath = os.Getenv("NOTES_PATH")
	if notesPath == "" {
		log.Fatal("NOTES_PATH environment variable must be set")
	}
	notesScratchFile = getEnvWithDefault("NOTES_SCRATCHFILE", "scratch")
	notesEditor = getEnvWithDefault("EDITOR", "vi")
	notesViewer = getEnvWithDefault("NOTES_VIEWER", "cat")

	// Ensure the NOTES_PATH directory exists
	ensureDirectoryExists(notesPath)
}

// appendToScratchFile handles appending piped input to the scratch file
func appendToScratchFile() {
	fullScratchPath := filepath.Join(notesPath, notesScratchFile+"."+notesFileType)
	file, err := os.OpenFile(fullScratchPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening scratch file: %v", err)
	}
	defer file.Close()

	// Read piped input and append to the scratch file
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		_, err := file.WriteString(scanner.Text() + "\n")
		if err != nil {
			log.Fatalf("Error writing to scratch file: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}

	fmt.Println("Text appended to file:", fullScratchPath)
}

// openScratchFileInEditor opens the scratch file in the editor for editing
func openScratchFileInEditor() {
	fullScratchPath := filepath.Join(notesPath, notesScratchFile+"."+notesFileType)

	fmt.Printf("Opening scratch file %s in editor: %s\n", fullScratchPath, notesEditor)
	editCmd := exec.Command(notesEditor, fullScratchPath)
	editCmd.Stdout = os.Stdout
	editCmd.Stderr = os.Stderr
	editCmd.Stdin = os.Stdin

	if err := editCmd.Run(); err != nil {
		log.Fatalf("Error launching editor: %v", err)
	}
}
