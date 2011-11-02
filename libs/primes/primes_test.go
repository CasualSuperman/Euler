package primes

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestPrimes(t *testing.T) {
	var limit int = 10000

	fmt.Println("Reading in results.")
	file, _ := os.Open("10000.txt")
	list := make([]uint, 0)
	char := []byte{}
	eof := false
	for !eof {
		single := []byte{0x00}
		_, done := file.Read(single)
		if done != nil {
			eof = true
		} else if string(single) == " " {
			num, _ := strconv.Atoui(string(char))
			list = append(list, num)
			char = []byte{}
		} else {
			char = append(char, single...)
		}
	}

	fmt.Println("Generating primes.")
	result := Primes(uint(limit))

	fmt.Println("Comparing results.")
	for i := 0; i < limit; i++ {
		found := false
		for a := 0; a < len(list) && !found; a++ {
			if list[a] == uint(i) {
				found = true
			}
		}
		if found != result.Value(uint(i)) {
			t.Error("WRONG")
			if found {
				fmt.Printf("Program incorrectly reported that %v is composite.\n", i)
			} else {
				fmt.Printf("Program incorrectly determined that %v is prime.\n", i)
			}
		}
	}
}
