package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	noFailHeader := http.Header{}
	noFailHeader.Add("Authorization", "ApiKey 123456789")
	malformedHeader := http.Header{}
	malformedHeader.Add("Authorization", "No Key")
	tests := map[string]struct {
		input         http.Header
		expected      string
		expectedError string
	}{
		"no fail": {
			input:    noFailHeader,
			expected: "123456789",
		},
		"empty header": {
			input:         http.Header{},
			expected:      "",
			expectedError: "no authorization header included",
		},
		"malformed header": {
			input:         malformedHeader,
			expected:      "",
			expectedError: "malformed authorization header",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := GetAPIKey(tc.input)
			diffResult := cmp.Diff(tc.expected, actual)
			if diffResult != "" {
				t.Fatal(diffResult)
			}
			if err != nil {
				diffErr := cmp.Diff(tc.expectedError, err.Error())
				if diffErr != "" {
					t.Fatal(diffErr)
				}
			}
		})
	}
}
