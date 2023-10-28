package pow

import "testing"

func TestPoWWithDefaultParams(t *testing.T) {
	challengeLength := DefaultChallengeLength
	requiredZeros := DefaultRequiredZeros

	challenge := GenerateChallenge(challengeLength)
	if len(challenge) != challengeLength {
		t.Fatalf("Expected challenge length %d, but got %d", challengeLength, len(challenge))
	}

	solution, err := SolveChallenge(challenge, requiredZeros)
	if err != nil {
		t.Fatalf("Failed to solve challenge: %v", err)
	}

	if !CheckPoW(challenge, solution, requiredZeros) {
		t.Fatalf("The solution did not pass the PoW check")
	}
}

func TestPoWWithOtherParams(t *testing.T) {
	challengeLength := 15
	requiredZeros := 6

	challenge := GenerateChallenge(challengeLength)
	if len(challenge) != challengeLength {
		t.Fatalf("Expected challenge length %d, but got %d", challengeLength, len(challenge))
	}

	solution, err := SolveChallenge(challenge, requiredZeros)
	if err != nil {
		t.Fatalf("Failed to solve challenge: %v", err)
	}

	if !CheckPoW(challenge, solution, requiredZeros) {
		t.Fatalf("The solution did not pass the PoW check")
	}
}

func BenchmarkSolveChallenge(b *testing.B) {
	requiredZeros := DefaultRequiredZeros
	challenge := []byte("qrPRkAyzqS")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := SolveChallenge(challenge, requiredZeros)
		if err != nil {
			b.Error("Failed to solve challenge")
		}
	}
}
