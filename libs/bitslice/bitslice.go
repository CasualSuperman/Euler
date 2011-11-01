package BitSlice

import "sync"

type BitSlice struct {
	arr   []byte
	locks []sync.RWMutex
	len   uint
}

func New(length uint) BitSlice {
	byteLen := length / 8
	if length%8 > 0 {
		newLen++
	}

	lockLen := byteLen / 8
	if byteLen%8 > 0 {
		lockLen++
	}

	return BitSlice{make([]byte, byteLen),
		make([]sync.RWMutex, lockLen),
		length}
}

/* Helper Methods */
func (b BitSlice) checkBounds(index uint) {
	if index >= b.len {
		panic("Index out of range.")
	}
}

func getIndexMask(i uint) (byteIndex, lockIndex uint, mask byte) {
	byteIndex = i / 8
	lockIndex = byteIndex / 8
	mask = 0x01 << (7 - (i % 8))
	return
}

/* Public Methods */
func (b BitSlice) Len() uint {
	return b.len
}

func (b BitSlice) Value(index uint) bool {
	b.checkBounds(index)

	bIndex, lIndex, mask := getIndexMask(index)

	defer func() {
		b.locks[lIndex].RUnlock()
	}()

	b.locks[lIndex].RLock()
	return (b.arr[index] & mask) != 0
}

func (b BitSlice) SetValue(index uint, value bool) {
	b.checkBounds(index)

	bIndex, lIndex, mask := getIndexMask(index)
	b.locks[lIndex].Lock()

	if !value {
		b.arr[index] &= ^mask // Example: ANDed with 11111011
	} else {
		b.arr[index] |= mask // Example: ORed with 00010000
	}

	b.locks[lIndex].Unlock()
}

func (b BitSlice) FlipValue(index uint) {
	b.checkBounds(index)

	bIndex, lIndex, mask := getIndexMask(index)
	b.locks[lIndex].Lock()

	b.arr[index] ^= mask

	b.locks[lIndex].Unlock()
}
