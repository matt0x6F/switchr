package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version displays the current version
var Version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays version",
	Long:  `Displays the version of the application`,
	Run:   versionCommand,
}

func versionCommand(cmd *cobra.Command, args []string) {
	if Version == "" {
		Version = "dev"
	}
	fmt.Printf("switchr version: %s\n", Version)
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
