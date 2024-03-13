package jwt

// import (
// 	"crypto/rand"
// 	"os"
// 	"regexp"
// 	"strings"
// )

// // HexRegexp is a regular expression to match hexadecimal characters.
// var HexRegexp = regexp.MustCompile(`^(?:0x)?[0-9a-fA-F]*$`)

// // JWTLength defines the length of the JWT byte array to be 32 bytes as
// // defined the Engine API specification.
// // https://github.com/ethereum/execution-apis/blob/main/src/engine/authentication.md
// //
// //nolint:lll
// const JWTLength = 32

// // Secret represents a JSON Web Token as a fixed-size byte array.
// type Secret [JWTLength]byte

// // NewFromFile reads the JWT secret from a file and returns it.
// // func NewFromFile(filepath string) (*Secret, error) {
// // 	data, err := os.ReadFile(filepath)
// // 	if err != nil {
// // 		// Return an error if the file cannot be read.
// // 		return nil, err
// // 	}
// // 	return NewFromHex(strings.TrimSpace(string(data)))
// // }

// // // NewFromHex creates a new JWT secret from a hexadecimal string.
// // func NewFromHex(hexStr string) (*Secret, error) {
// // 	// Ensure the hex string contains only hexadecimal characters.
// // 	if !HexRegexp.MatchString(hexStr) {
// // 		return nil, ErrContainsIllegalCharacter
// // 	}

// // 	// Convert the hex string to a byte array.
// // 	bz, err := bytes.De
// // 	if bz == nil || len(bz) != EthereumJWTLength {
// // 		return nil, ErrLengthMismatch
// // 	}
// // 	s := Secret(bz)
// // 	return &s, nil
// // }

// func NewFromHex(hexStr string) (*Secret, error) {
// 	// Ensure the hex string contains only hexadecimal characters.
// 	if !HexRegexp.MatchString(hexStr) {
// 		return nil, ErrInvalidHex
// 	}

// 	// Convert the hex string to a byte array.
// 	bz, err := hexutil.Decode(hexStr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if bz == nil || len(bz) != JWTLength {
// 		return nil, ErrLengthMismatch
// 	}
// 	s := Secret(bz)
// 	return &s, nil
// }

// // NewRandom creates a new random JWT secret.
// func NewRandom() (*Secret, error) {
// 	secret := make([]byte, JWTLength)
// 	// We don't need to check n since:
// 	// n == len(b) if and only if err == nil.
// 	_, err := rand.Read(secret)
// 	if err != nil {
// 		return nil, err
// 	}
// // 	return NewFromHex(hexutil.Encode(secret))
// }

// // String returns the JWT secret as a string with the first 8 characters
// // visible and the rest masked out for security.
// func (s *Secret) String() string {
// 	// secret := hexutil.Encode(s[:])
// 	// return secret[:8] + strings.Repeat("*", len(secret[8:]))
// }

// // Hex returns the JWT secret as a hexadecimal string.
// func (s *Secret) Hex() string {
// 	// return hexutil.Encode(s[:])
// }

// // Bytes returns the JWT secret as a byte array.
// func (s *Secret) Bytes() []byte {
// 	return s[:]
// }
