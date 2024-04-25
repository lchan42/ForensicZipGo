package parser

import (
	"crypto"
	"testing"
)

func TestParseHashAlg(t *testing.T) {
	tests := []struct {
		hashName      string
		expectedHash  crypto.Hash
		expectedError bool
	}{
		{"MD5", crypto.MD5, false},
		{"SHA256", crypto.SHA256, false},
		{"SHA384", crypto.SHA384, false},
		{"NonExistentHash", 0, true},
	}

	for _, test := range tests {
		result, err := ParseHashAlg(test.hashName)

		if test.expectedError && err == nil {
			t.Errorf("Expected error but got nil for hashName: %s", test.hashName)
		}

		if !test.expectedError && err != nil {
			t.Errorf("Unexpected error for hashName: %s - Error: %v", test.hashName, err)
		}

		if result != test.expectedHash {
			t.Errorf("Unexpected result for hashName: %s - Expected: %v, Got: %v", test.hashName, test.expectedHash, result)
		}
	}
}
