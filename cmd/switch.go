package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/mattouille/switchr/config"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// Might be good to use Errof (https://golang.org/pkg/errors/#example_New_errorf)

var (
	errMissingArg      = errors.New("Missing user argument")
	errProfileNotFound = errors.New("No profile found")
	errHomeDir         = errors.New("Determining home directory")
	errOpenFile        = errors.New("Opening file")
	errCreateFile      = errors.New("Creating file")
	errCopyFile        = errors.New("Copying file")
	errSyncFile        = errors.New("Syncing file")
	errExecCmd         = errors.New("Executing command")
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch user",
	Long:  `Switch the active user`,
	Run:   switchCommand,
}

func init() {
	RootCmd.AddCommand(switchCmd)
}

func switchCommand(cmd *cobra.Command, args []string) {
	if err := processArgs(args); err != nil {
		log.Fatalf("Error %s", err)
	}
}

func processArgs(args []string) error {
	if len(args) == 0 {
		return errMissingArg
	}

	if err := findProfile(args[0], configuration.Profiles); err != nil {
		return err
	}

	return nil
}

func findProfile(user string, profiles []config.ProfileConfiguration) error {
	for _, profile := range profiles {
		// Look for the profile by email address
		if profile.Email != user {
			continue
		}

		// Profile found
		if err := switchUser(profile); err != nil {
			// Error occurred during switching
			return err
		}

		// Profile found, no errors
		return nil
	}

	return errProfileNotFound
}

func switchUser(profile config.ProfileConfiguration) error {
	home, err := homedir.Dir()
	if err != nil {
		return errHomeDir
	}

	if err := moveKeys(home, profile); err != nil {
		return err
	}

	fmt.Printf("Setting Git name to %s and email to %s\n", profile.Name, profile.Email)

	if err := setGit("user.name", profile.Name); err != nil {
		return err
	}

	if err := setGit("user.email", profile.Email); err != nil {
		return err
	}

	return nil
}

func setGit(arg string, value string) error {
	cmd := "git"
	args := []string{"config", "--global", arg, value}

	if err := exec.Command(cmd, args...).Run(); err != nil {
		return errExecCmd
	}

	return nil
}

func moveKeys(home string, profile config.ProfileConfiguration) error {
	hostPrivateKey := home + "/.ssh/id_rsa"
	hostPublicKey := home + "/.ssh/id_rsa.pub"
	userPrivateKey := home + "/.ssh/" + profile.Key
	userPublicKey := home + "/.ssh/" + profile.Key + ".pub"

	fmt.Printf(
		"Moving ssh key %s/.ssh/%s, %s/.ssh/%s.pub\n",
		home, profile.Key, home, profile.Key,
	)

	if err := copyFile(userPrivateKey, hostPrivateKey); err != nil {
		return err
	}
	if err := copyFile(userPublicKey, hostPublicKey); err != nil {
		return err
	}

	return nil
}

func copyFile(source string, destination string) error {
	srcFile, err := os.Open(source)
	if err != nil {
		return errOpenFile
	}
	defer srcFile.Close()

	// Create if the file doesn't exist
	destFile, err := os.Create(destination)
	if err != nil {
		return errCreateFile
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return errCopyFile
	}

	err = destFile.Sync()
	if err != nil {
		return errSyncFile
	}

	return nil
}
