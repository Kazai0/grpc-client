package utils

import (
	"encoding/binary"
	"math/rand"
	"time"
)

//function to create number random in bytes
func GenerateRandomBytes() ([]byte, error) {
	// Initialize the random number generator with a unique seed
	rand.Seed(time.Now().UnixNano())

	// Generate a random number (int64, for example)
	randomNumber := rand.Int63()

	// Create a byte slice with 8 bytes to store the number
	randomBytes := make([]byte, 8)

	// Convert the number to bytes
	binary.BigEndian.PutUint64(randomBytes, uint64(randomNumber))

	return randomBytes, nil
}
