package main

func main() {
	total := 1
	rounds := (1001 - 1) / 2 // Number of sets of four numbers
	sideLength := 1
	lastMax := 1
	for i := 0; i < rounds; i++ {
		sideLength += 2
		for j := 0; j < 4 ; j++ {
			lastMax += sideLength - 1
			total += lastMax
		}
	}
	println (total)
}
