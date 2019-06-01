package main

import "../libs/primes"
import "fmt"
import "math"

func main() {
	max := formula{}
	candidates := primes.Primes(100000)
	for i := -1000; i < 1000; i++ {
		if candidates.Value(uint(math.Abs(float64(i)))) {
			// j must be larger than i to make n=0 positive
			for j := i + 1; j < 1000; j++ {
				if candidates.Value(uint(math.Abs(float64(j)))) {
					k := 0
					for candidates.Value(uint(math.Abs(float64(k * k + i * k + j)))) {
						k++
					}
					if max.count < k {
						max = formula{i, j, k}
					}
				}
			}
		}
	}
	fmt.Printf("n²+%dn+%d (%d)\n", max.a, max.b, max.count)
	fmt.Printf("%d\n", max.a * max.b)
}

type formula struct {
	a, b int
	count int
}

