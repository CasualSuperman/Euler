package main

import bs "../libs/bitslice"
import "../libs/primes"
//import "fmt"

func main() {
	found := 0
	var total uint = 0
	candidates := primes.Primes(1000000)
	var i uint = 11
	for found < 11 {
		if candidates.Value(i) {
			if isRightTruncPrime(i, candidates) && isLeftTruncPrime(i, candidates) {
				println(i)
				found++
				total += i
			}
		}
		i++
	}
	println("")
	println(total)
}

// There is an optimization to be made where if this method fails, we can skip primes with the same prefix,
// but this solution is fast enough without it
func isLeftTruncPrime(prime uint, primes *bs.BitSlice) bool {
	var place uint = 10
	for place * 10 < prime {
		place *= 10
	}
	//fmt.Println("Left: ", prime)
	for prime > 10 {
		prime %= place
		place /= 10
		//fmt.Println("\t", prime)
		if !primes.Value(prime) {
			return false
		}
	}
	return true
}

func isRightTruncPrime(prime uint, primes *bs.BitSlice) bool {
	//fmt.Println("Right: ", prime)
	for prime > 10 {
		prime /= 10
		//fmt.Println("\t", prime)
		if !primes.Value(prime) {
			return false
		}
	}
	return true
}
