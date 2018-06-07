package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all users",
	Long:  `Lists all users in switchr`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, profile := range configuration.Profiles {
			fmt.Printf("%s [%s]\n", profile.Name, profile.Email)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
