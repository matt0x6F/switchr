package cmd

import "testing"

func TestGetVersion(t *testing.T) {
	testCases := []struct {
		version string
		want    string
	}{
		{"", "dev"},
		{"1.0", "1.0"},
	}
	for _, tc := range testCases {
		version := getVersion(tc.version)
		if version != tc.want {
			t.Errorf("Version was incorrect, got %s, want %s", version, tc.want)
		}
	}
}
