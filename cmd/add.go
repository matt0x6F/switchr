package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a user",
	Long:  `Adds a user to switchr`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Printf("Adding user: %s\n", args[0])
		} else {
			fmt.Println("Error: you must provide a user")
			fmt.Println("Usage: switchr add [user]")
			os.Exit(2)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
