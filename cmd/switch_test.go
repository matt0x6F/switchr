package cmd

import "testing"

func TestProcessArgs(t *testing.T) {
	// https://stackoverflow.com/questions/40615641/testing-os-exit-scenarios-in-go-with-coverage-information-coveralls-io-goverall/40801733
	// Save current function and restore at the end:
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit
	processArgs([]string{})

	if exp := 2; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}
}
