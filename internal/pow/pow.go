package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	DefaultChallengeLength = 10
	DefaultRequiredZeros   = 4
)

func GenerateChallenge(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = chars[r.Intn(len(chars))]
	}
	return string(b)
}

func CheckPoW(challenge, solution string, requiredZeros int) bool {
	data := challenge + solution
	hash := sha256.Sum256([]byte(data))
	return strings.HasPrefix(hex.EncodeToString(hash[:]), strings.Repeat("0", requiredZeros))
}

func SolveChallenge(challenge string, requiredZeros int) string {
	for i := 0; ; i++ {
		solution := fmt.Sprintf("%d", i)
		data := challenge + solution
		hash := sha256.Sum256([]byte(data))
		if strings.HasPrefix(hex.EncodeToString(hash[:]), strings.Repeat("0", requiredZeros)) {
			return solution
		}
	}
}
