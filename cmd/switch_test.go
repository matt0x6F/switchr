package cmd

import "testing"

func TestProcessArgs(t *testing.T) {
	testCases := []struct {
		string
		want string
	}{
		{"", "dev"},
		{"1.0", "1.0"},
	}
}
