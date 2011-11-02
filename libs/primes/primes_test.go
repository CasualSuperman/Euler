package primes

import "testing"

func TestPrimes(t *testing.T) {
	var limit int = 20
	primes := map[int]bool {
		2: true,
		3: true,
		5: true,
		7: true,
		11: true,
		13: true,
		17: true,
		19: true}
	result := Primes(uint(limit))
	for i := 0; i < limit; i++ {
		if ! primes[i] == result.Value(uint(i)) {
			t.Error("WRONG")
		}
	}
}
