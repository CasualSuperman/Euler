package primes

import bs "../bitslice/bitslice"

func Primes(limit uint) *bs.BitSlice {
	length := int(limit / 8)
	if limit%8 > 0 {
		length++
	}

	list := make([]byte, uint(length))

	// Initialize for values 2, 3, 5, and 7 already run.
	// 0123456789 10 11 12 13 14 15 16 17...
	// 0011010100  0  1  0  1  0  0  0  1
	list[0] = 0x35
	for i := 1; i < length; i += 3 {
		list[i] = 0x14
	}
	for i := 2; i < length; i += 3 {
		list[i] = 0x51
	}
	for i := 3; i < length; i += 3 {
		list[i] = 0x45
	}

	// smallestGuaranteedPrime
	// The largest number guaranteed to be prime at this point.
	// This is based off the principle that in an iteration of the
	//     sieve of eratosthenes for number I, all values less than
	//     I^2 are final.
	// Our array has been initialized for i = 7
	var bigGPrime uint  = 49
	primes := bs.New(limit)
	primes.Arr = list
	return primes
}
