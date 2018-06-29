package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version displays version and is imported at build time
var Version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays version",
	Long:  `Displays the version of the application`,
	Run:   versionCommand,
}

func versionCommand(cmd *cobra.Command, args []string) {
	fmt.Printf("switchr version: %s\n", getVersion(Version))
}

func getVersion(version string) string {
	if version == "" {
		version = "dev"
	}

	return version
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
