package parser

import (
	"fmt"
	"crypto/md5"                      // for MD5
	"crypto/sha1"                     // for SHA1
	"crypto/sha256"                   // for SHA224 and SHA256
	"crypto/sha512"                   // for SHA384, SHA512, SHA512_224, and SHA512_256
	"golang.org/x/crypto/sha3"        // for SHA3_224, SHA3_256, SHA3_384, SHA3_512
	"golang.org/x/crypto/blake2b"     // for BLAKE2b_256, BLAKE2b_384, BLAKE2b_512
	"golang.org/x/crypto/blake2s"     // for BLAKE2s_256
	"hash"
)


func ParseHashAlg(hashName string) (func() hash.Hash, error){

	var hashFunctions = map[string]func() hash.Hash{
		"MD5"				:	md5.New,
		"SHA1"				:	sha1.New,
		"SHA224"			:	sha256.New224,
		"SHA256"			:	sha256.New,
		"SHA384"			:	sha512.New384,
		"SHA512"			:	sha512.New,
		"SHA512_224"		:	sha512.New512_224,
		"SHA512_256"		:	sha512.New512_256,
		"SHA3_224"			:	sha3.New224,
		"SHA3_256"			:	sha3.New256,
		"SHA3_384"			:	sha3.New384,
		"SHA3_512"			:	sha3.New512,
		"BLAKE2s_256"		:	func() hash.Hash { h, _ := blake2s.New256(nil); return h },
		"BLAKE2b_256"		:	func() hash.Hash { h, _ := blake2b.New256(nil); return h },
		"BLAKE2b_384"		:	func() hash.Hash { h, _ := blake2b.New384(nil); return h },
		"BLAKE2b_512"		:	func() hash.Hash { h, _ := blake2b.New512(nil); return h },
	}

	if value, ok := hashFunctions[hashName]; !ok{
		return nil, fmt.Errorf(
			`unvalide hash parameter.
			Available has are :
			- MD4
			- MD5
			- SHA1
			- SHA224
			- SHA256
			- SHA384
			- SHA512
			- SHA512_224
			- SHA512_256
			- SHA3_224
			- SHA3_256
			- SHA3_384
			- SHA3_512
		`)
			// add Blake2 hash accordingly
			// - BLAKE2s_256
			// - BLAKE2b_256
			// - BLAKE2b_384
			// - BLAKE2b_512

	} else {
		// fmt.Printf("hash function = %v\n", value)
		return value, nil
	}
}

