package parser

import (
	"hash"
	"testing"
)

// TestParseHashAlgValid checks if the ParseHashAlg function correctly returns a non-nil hash function for valid inputs.
func TestParseHashAlgValid(t *testing.T) {
	validHashes := []string{
		"MD5", "SHA1", "SHA224", "SHA256", "SHA384", "SHA512",
		"SHA512_224", "SHA512_256", "SHA3_224", "SHA3_256", "SHA3_384", "SHA3_512",
		"BLAKE2s_256", "BLAKE2b_256", "BLAKE2b_384", "BLAKE2b_512",
	}

	for _, hashName := range validHashes {
		hasherFunc, err := ParseHashAlg(hashName)
		if err != nil {
			t.Errorf("ParseHashAlg(%q) returned an error: %v", hashName, err)
		}
		if hasherFunc == nil {
			t.Errorf("ParseHashAlg(%q) returned nil function", hashName)
		} else {
			// Optionally check if returned func is a valid hash.Hash by calling it and checking for a non-nil result
			if h := hasherFunc(); h == nil || !checkHasher(h) {
				t.Errorf("ParseHashAlg(%q) returned an invalid hash function", hashName)
			}
		}
	}
}

// TestParseHashAlgInvalid tests if the ParseHashAlg function correctly handles invalid inputs.
func TestParseHashAlgInvalid(t *testing.T) {
	invalidHashes := []string{
		"invalid", "SHA256-invalid", "", "123", "SHA3999",
	}

	for _, hashName := range invalidHashes {
		hasherFunc, err := ParseHashAlg(hashName)
		if err == nil {
			t.Errorf("ParseHashAlg(%q) expected an error but got none", hashName)
		}
		if hasherFunc != nil {
			t.Errorf("ParseHashAlg(%q) expected nil function but got one", hashName)
		}
	}
}

// checkHasher is a helper function to ensure the hash.Hash actually operates
func checkHasher(h hash.Hash) bool {
	// Attempt to write something to the hasher and check sum it
	_, err := h.Write([]byte("test"))
	if err != nil {
		return false
	}
	sum := h.Sum(nil)
	return len(sum) > 0
}
