package main

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
	result := char{letter, -1}
	switch letter {
		case 'M':
			result.value = 1000
		case 'D':
			result.value = 500
		case 'C':
			result.value = 100
		case 'L':
			result.value = 50
		case 'X':
			result.value = 10
		case 'V':
			result.value = 5
		case 'I':
			result.value = 1
		default:
			panic("Invalid letter.")
	}
	return result
}

func (c char) Compare(other char) int {
	if c.value == other.value {
		return 0
	} else if c.value > other.value {
		return 1
	}
	return -1
}
