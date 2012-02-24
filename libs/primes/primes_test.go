package primes

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestPrimes(t *testing.T) {
	var limit int = 1047

	fmt.Println("Reading in results.")
	file, _ := os.Open("10000.txt")
	list := make(map[uint]bool)
	char := []byte{}
	eof := false
	for !eof {
		single := []byte{0x00}
		_, done := file.Read(single)
		if done != nil {
			eof = true
		} else if string(single) == " " {
			num, _ := strconv.ParseUint(string(char), 10, 0)
			list[num] = true
			char = []byte{}
		} else {
			char = append(char, single...)
		}
	}

	fmt.Println("Generating primes.")
	result := Primes(uint(limit))

	fmt.Println("Comparing results.")
	for i := 0; i < limit; i++ {
		if list[uint(i)] != result.Value(uint(i)) {
			if list[uint(i)] {
				t.Errorf("Program incorrectly reported that %v is composite.\n", i)
			} else {
				t.Errorf("Program incorrectly determined that %v is prime.\n", i)
			}
		}
	}
}

func BenchmarkPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Primes(1000000000)
	}
}
