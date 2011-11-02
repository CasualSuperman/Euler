package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var numerals []string

func main() {
	data, _ := ioutil.ReadFile("roman.txt")
	numerals = strings.Split(string(data), "\n")
	total := 0
	for _, numeral := range numerals {
		total += len(numeral)
		total -= len(findMinimalRepresentation(numeral))
	}
	fmt.Println(total)
}

func findMinimalRepresentation(s string) string {
	stack := []char{}
	for _, val := range s {
		stack = append(stack, NewChar(uint8(val)))
	}
	return ""
}
