package main

import (
	"io/ioutil"
)

const (
	WIDTH  = 80
	HEIGHT = 80
)

func main() {
	
}

func traverse(maze [WIDTH][HEIGHT]int) (sum int) {
	sum = maze[0][0]
	path := make([][2]int)
}
