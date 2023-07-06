package password

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"io/ioutil"

	"golang.org/x/crypto/bcrypt"
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

// Salt a password using SHA-512 hashing algorithm
func SaltPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the salted hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 salted password
	var saltedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the salted password to a hex string
	var saltedPasswordHex = hex.EncodeToString(saltedPasswordBytes)

	return saltedPasswordHex
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ReadResponseBody(body io.ReadCloser) []byte {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	return bodyBytes
}
