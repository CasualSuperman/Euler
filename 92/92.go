package main

func main() {
	num := 0

	for i := 1; i < 10000000; i++ {
		chainTermination := doChain(i)
		if chainTermination == 89 {
			num++
		}
	}

	println(num)
}

func doChain(i int) int {
	for i != 1 && i != 89 {
		i = nextInChain(i)
	}

	return i
}

func nextInChain(i int) int {
	total := 0

	for i > 0 {
		j := i%10
		total += j*j
		i /= 10
	}

	return total
}
