package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mattouille/switchr/config"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var (
	cfgFile       string
	configuration config.Configuration
	RootCmd       = &cobra.Command{
		Use:   "switchr",
		Short: "Switchr is a profile switcher for pair programming",
		Long:  `Switchr is a profile switcher for pair programming in Linux on shared workstation`,
		Run: func(cmd *cobra.Command, args []string) {
			//
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
}

// Execute the command or exit with an error
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// initConfig reads in a config file and ENV variables if set.
func initConfig() {

	// Find home directory
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".switchr"
	viper.AddConfigPath(home)
	viper.SetConfigName(".switchr")

	viper.AutomaticEnv() // set environment variables

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
	// Unmarshal to configuration
	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
