package primes

import (
	bs "../bitslice/bitslice"
)

const MAX_CONCURRENT = 4

func Primes(limit uint) *bs.BitSlice {
	length := int(limit / 8)
	if limit%8 > 0 {
		length++
	}

	list := make([]byte, uint(length))

	// Initialize for values 2, 3, and 5 already run.
	// Avoids the nasty small loops
	// 0123456789 10 11 12 13 14 15 16 17...
	// 0011010100  0  1  0  1  0  0  0  1
	list[0] = 0x35 // Special case, 2 is prime
	for i := 1; i < length; i += 3 {
		list[i] = 0x14
	}
	for i := 2; i < length; i += 3 {
		list[i] = 0x51
	}
	for i := 3; i < length; i += 3 {
		list[i] = 0x45
	}

	primes := bs.New(limit)
	primes.Arr = list
	generate(primes, limit)
	return primes
}

func generate(primes *bs.BitSlice, limit uint) {
	// biggestGuaranteedPrime
	// The largest number guaranteed to be determined at this point.
	// This is based off the principle that in an iteration of the
	//     sieve of eratosthenes for number I, all values less than
	//     I^2 are final.
	// Our array has been initialized for i = 5
	var bigGPrime uint = 25
	// The last number set to run
	var lastGen uint = 7

	var generating = make(map[uint]bool, MAX_CONCURRENT)
	var done = make(chan uint)
	// Loop until we reach the end.
	for bigGPrime < limit {
		// Launch sieves
		for len(generating) < MAX_CONCURRENT && lastGen < bigGPrime {
			// Skip values that aren't prime
			for !primes.Value(lastGen) {
				// Stop if we go too far
				if lastGen < bigGPrime {
					break
				}
				lastGen += 6 // Skip multiples of 2 and 3
			}
			// If we didn't go too far, we find a prime that needs sieving
			if lastGen <= bigGPrime {
				generating[lastGen] = true
				go run(primes, lastGen, limit, done)
			}
		}
		// If we're stuck, either due to surpassing our max threads or
		// passing our biggest known value
		if len(generating) >= MAX_CONCURRENT || lastGen >= bigGPrime {
			// wait.
			val := <-done
			generating[val] = false, false // Remove it from the list

			// Find the new largest known value (This might not change)
			smallest := limit
			for val, _ := range generating {
				if val < smallest {
					smallest = val
				}
			}
			bigGPrime = uint(smallest * smallest) - 1
		}
		// Repeat
	}
}

func run(slice *bs.BitSlice, val, max uint, done chan uint) {
	start := val
	val = start * start
	for val <= max {
		val += start
		if slice.Value(val) {
			slice.SetValue(val, false)
		}
	}
	done <- start
}
