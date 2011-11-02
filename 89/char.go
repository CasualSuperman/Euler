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
	// Map from letter to value
	vals :=  map[byte] int {
		'M': 1000,
		'D':  500,
		'C':  100,
		'L':   50,
		'X':   10,
		'V':    5,
		'I':    1}

	result := char{letter, 0x111E6A1} // Illegal
	switch letter {
		case 'M','D','C','L','X','V','I':
			result.value = vals[letter]
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
