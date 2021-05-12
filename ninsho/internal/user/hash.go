package user

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type hashParams struct {
	memory      uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func NewHashParams() *hashParams {
	return &hashParams{
		memory:      64 * 1024,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
}

func GenerateSalt(config *hashParams) ([]byte, error) {
	salt := make([]byte, config.saltLength)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func Hash(config *hashParams, password string) ([]byte, error) {
	salt, err := GenerateSalt(config)
	if err != nil {
		return nil, nil
	}
	key := argon2.IDKey([]byte(password), salt, 1, config.memory, config.parallelism, config.keyLength)
	return key, nil
}

func Verify(config *hashParams, hash string, plain string) (bool, error) {
	hashParts := strings.Split(hash, "$")

	_, err := fmt.Sscanf(hashParts[3], "m=%d,t=%d,p=%d", &config.memory, 1, &config.parallelism)

	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(hashParts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(hashParts[5])
	if err != nil {
		return false, err
	}

	hashToCompare := argon2.IDKey([]byte(hash), salt, 1, config.memory, config.parallelism, uint32(len(decodedHash)))

	return subtle.ConstantTimeCompare(decodedHash, hashToCompare) == 1, nil
}
