package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"

	"github.com/xdg-go/pbkdf2"
)

// Define salt size
const SaltSize = 32

// GenerateRandomSalt generates a random salt byte slice
func GenerateRandomSalt(saltSize int) string {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(salt[:])
}

func SaltAndHashPassword(password string, salt string) string {
	dk := pbkdf2.Key([]byte(password), []byte(salt), 4096, 32, sha1.New)
	return hex.EncodeToString(dk[:])
}

func CheckPassword(password string, salt string, hash string) bool {
	return SaltAndHashPassword(password, salt) == hash
}

func ReadResponseBody(body io.ReadCloser) []byte {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	return bodyBytes
}
