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
		byteLen++
	}

	lockLen := byteLen / 8
	if byteLen%8 > 0 {
		lockLen++
	}

	return BitSlice{make([]byte, byteLen),
		make([]sync.RWMutex, lockLen),
		length}
}

func Quick(data []byte) BitSlice {
	lockLen := len(data) / 8
	if len(data)%8 > 0 {
		lockLen++
	}
	return BitSlice{data, make([]sync.RWMutex, lockLen), uint(len(data))}
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
	return (b.arr[bIndex] & mask) != 0
}

func (b BitSlice) SetValue(index uint, value bool) {
	b.checkBounds(index)

	bIndex, lIndex, mask := getIndexMask(index)
	b.locks[lIndex].Lock()

	if !value {
		b.arr[bIndex] &= ^mask // Example: ANDed with 11111011
	} else {
		b.arr[bIndex] |= mask // Example: ORed with 00010000
	}

	b.locks[lIndex].Unlock()
}

func (b BitSlice) FlipValue(index uint) {
	b.checkBounds(index)

	bIndex, lIndex, mask := getIndexMask(index)
	b.locks[lIndex].Lock()

	b.arr[bIndex] ^= mask

	b.locks[lIndex].Unlock()
}
