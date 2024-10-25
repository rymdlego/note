package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "v0.1.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Note version: %v\n", Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
