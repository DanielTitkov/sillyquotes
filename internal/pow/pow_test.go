package pow

import "testing"

func TestPoWWithDefaultParams(t *testing.T) {
	challengeLength := DefaultChallengeLength
	requiredZeros := DefaultRequiredZeros

	challenge := GenerateChallenge(challengeLength)
	if len(challenge) != challengeLength {
		t.Fatalf("Expected challenge length %d, but got %d", challengeLength, len(challenge))
	}

	solution := SolveChallenge(challenge, requiredZeros)

	if !CheckPoW(challenge, solution, requiredZeros) {
		t.Fatalf("The solution did not pass the PoW check")
	}
}

func TestPoWWithOtherParams(t *testing.T) {
	challengeLength := 12
	requiredZeros := 5

	challenge := GenerateChallenge(challengeLength)
	if len(challenge) != challengeLength {
		t.Fatalf("Expected challenge length %d, but got %d", challengeLength, len(challenge))
	}

	solution := SolveChallenge(challenge, requiredZeros)

	if !CheckPoW(challenge, solution, requiredZeros) {
		t.Fatalf("The solution did not pass the PoW check")
	}
}
