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

type char struct {
	letter byte
	value  int
}

var (
	M = char{'M', 1000}
	D = char{'D',  500}
	C = char{'C',  100}
	L = char{'L',   50}
	X = char{'X',   10}
	V = char{'V',    5}
	I = char{'I',    1}
)

func NewChar(letter uint8) char {
	switch letter {
		case 'M':
			return M
		case 'D':
			return D
		case 'C':
			return C
		case 'L':
			return L
		case 'X':
			return X
		case 'V':
			return V
		case 'I':
			return I
	}
	panic("Unreachable")
}

func (c char) Compare(other char) int {
	if c.value == other.value {
		return 0
	} else if c.value > other.value {
		return 1
	}
	return -1
}

func findMinimalRepresentation(s string) string {
	stack := []char{}
	for _, val := range s {
		stack = append(stack, NewChar(uint8(val)))
	}
	return ""
}
