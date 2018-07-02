package cmd

import (
	"testing"

	"github.com/mattouille/switchr/config"
)

func TestProcessArgs(t *testing.T) {
	testCases := []struct {
		args []string
		want error
	}{
		{[]string{}, errMissingArg},
	}

	for _, tc := range testCases {
		test := processArgs(tc.args)
		if test != tc.want {
			t.Errorf("Error message was incorrect, got %s, want %s", test, tc.want)
		}
	}
}

func TestFindProfile(t *testing.T) {
	testProfile := &config.ProfileConfiguration{
		Name:  "Test User",
		Email: "test@test.com",
		Key:   "test",
	}

	configuration := &config.Configuration{
		Profiles: []config.ProfileConfiguration{*testProfile},
	}

	testCases := []struct {
		email string
		want  error
	}{
		{testProfile.Email, errOpenFile},
	}

	for _, tc := range testCases {
		test := findProfile(tc.email, configuration.Profiles)
		if test != tc.want {
			t.Errorf("Error message was incorrect, got %s, want %s", test, tc.want)
		}
	}
}
