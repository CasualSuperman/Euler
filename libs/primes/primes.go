package primes

import (
	bs "../bitslice"
	//"fmt"
	"runtime"
)

const MAX_CONCURRENT = 8

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
	runtime.GOMAXPROCS(MAX_CONCURRENT)
	generate(primes, limit)
	runtime.GOMAXPROCS(1)
	return primes
}

var alt = false

func increment(i uint) uint {
	switch {
	case alt:
		i += 2
		fallthrough
	default:
		i += 2
	}
	alt = !alt
	return i
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
	var lastGen uint = 5

	var generating = make(map[uint]bool, MAX_CONCURRENT)
	var generated = make(map[uint]bool)
	var done = make(chan uint)
	// Loop until we reach the end.
	for bigGPrime <= limit {
		// Launch sieves
		for len(generating) < MAX_CONCURRENT && lastGen <= bigGPrime {
			// Skip values that aren't prime
			for !primes.Value(lastGen) {
				//fmt.Printf("Checking %v\n", lastGen)
				// Stop if we go too far
				if lastGen > bigGPrime {
					break
				}
				lastGen = increment(lastGen)
			}
			// If we didn't go too far, we found a prime that needs sieving
			if lastGen <= bigGPrime {
				generating[lastGen] = true
				//fmt.Printf("Number of runs still out: %v\n", len(generating))
				go run(primes, lastGen, limit, done)
				lastGen = increment(lastGen)
			}
		}
		// If we're stuck, either due to surpassing our max threads or
		// passing our biggest known value
		if len(generating) >= MAX_CONCURRENT || lastGen > bigGPrime {
			mostRecent := <-done
			delete(generating, mostRecent) // Remove it from the list
			generated[mostRecent] = true
			// Find the new largest known value (This might not change)
			if len(generating) == 0 {
				biggest := uint(0)
				for val, _ := range generated {
					if val > biggest {
						biggest = val
					}
				}
				bigGPrime = uint(biggest * biggest)
			} else {
				smallest := mostRecent
				for val, _ := range generating {
					if val < smallest {
						//fmt.Printf("New smallest routine out: %v\n", val)
						smallest = val
					}
				}
				bigGPrime = uint(smallest * smallest)
			}
			// fmt.Printf("New bigGPrime is %v.\n", bigGPrime)
		}
		// Repeat
	}
	for len(generating) > 0 {
		delete(generating, <-done)
	}
}

func run(slice *bs.BitSlice, val, max uint, done chan uint) {
	//fmt.Printf("Launching sieve of %v.\n", val)
	start := val
	val = start * start
	for val < max {
		// fmt.Printf("Sieve of %v reached val %v.\n", start, val)
		go slice.SetValue(val, false)
		val += start
	}
	done <- start
}
