package pow

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"math/rand"
)

const (
	DefaultChallengeLength = 12
	DefaultRequiredZeros   = 5
	chars                  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GenerateChallenge(length int) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}

	return []byte(b)
}

func CheckPoW(challenge []byte, solution uint32, requiredZeros int) bool {
	solutionBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(solutionBytes, solution)
	data := append(challenge, solutionBytes...)
	hash := sha256.Sum256(data)

	requiredBytes := requiredZeros / 2
	for i := 0; i < requiredBytes; i++ {
		if hash[i] != 0 {
			return false
		}
	}

	if requiredZeros%2 != 0 {
		if hash[requiredBytes]&0xF0 != 0 {
			return false
		}
	}

	return true
}

func SolveChallenge(challenge []byte, requiredZeros int) (uint32, error) {
	var i uint32
	maxUint32 := uint32(0xFFFFFFFF)
	for i = 0; i < maxUint32; i++ {
		if CheckPoW(challenge, i, requiredZeros) {
			return i, nil
		}
	}

	return 0, errors.New("failed to find a solution")
}
