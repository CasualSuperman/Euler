package BitSlice

type BitSlice struct {
	arr []byte
	len uint
}

func New(length uint) BitSlice {
	newLen := length / 8
	if length % 8 > 0 {
		newLen++
	}
	return BitSlice{make([]byte, newLen), length}
}

/* Helper Methods */
func (b BitSlice) checkBounds(index uint) {
	if index >= b.len {
		panic("Index out of range.")
	}
}

func getIndexMask(i uint) (index uint, mask byte) {
	index = i/8
	mask  = 0x01 << (7 - (i % 8))
	return
}

/* Public Methods */
func (b BitSlice) Len() uint {
	return b.len
}

func (b BitSlice) Value(index uint) bool {
	b.checkBounds(index)

	index, mask := getIndexMask(index)
	return (b.arr[index] & mask) != 0
}

func (b BitSlice) SetValue(index uint, value bool) {
	b.checkBounds(index)

	index, mask := getIndexMask(index)
	if !value {
		b.arr[index] &= ^mask  // Example: ANDed with 11111011
	} else {
		b.arr[index] |= mask // Example: ORed with 00010000
	}
}

func (b BitSlice) FlipValue(index uint) {
	b.checkBounds(index)

	index, mask := getIndexMask(index)
	b.arr[index] ^= mask
}
