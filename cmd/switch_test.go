package cmd

import (
	"testing"
)

func TestProcessArgs(t *testing.T) {
	testCases := []struct {
		args []string
		want error
	}{
		{[]string{}, errMissingArg},
		{[]string{"profile-does-not-exist"}, errProfileNotFound},
	}

	for _, tc := range testCases {
		test := processArgs(tc.args)
		if test != tc.want {
			t.Errorf("Error message was incorrect, got %s, want %s", test, tc.want)
		}
	}
}
