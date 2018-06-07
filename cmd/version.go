package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays version",
	Long:  `Displays the version of the application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("switchr (v0.0.1)")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
