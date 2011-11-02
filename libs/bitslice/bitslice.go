package BitSlice

import "sync"

type BitSlice struct {
	Arr   []byte
	locks []sync.RWMutex
	len   uint
}

func New(length uint) *BitSlice {
	slice := new(BitSlice)
	byteLen := length / 8
	if length%8 > 0 {
		byteLen++
	}

	lockLen := byteLen / 64
	if byteLen%64 > 0 {
		lockLen++
	}

	slice.Arr = make([]byte, byteLen)
	slice.locks = make([]sync.RWMutex, lockLen)
	slice.len = length
	return slice
}

/* Helper Methods */
func (b BitSlice) checkBounds(index uint) {
	if index >= b.len {
		panic("Index out of range.")
	}
}

func getIndexMask(i uint) (byteIndex, lockIndex uint, mask byte) {
	byteIndex = i / 8
	lockIndex = byteIndex / 64
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
	return (b.Arr[bIndex] & mask) != 0
}

func (b BitSlice) SetValue(index uint, value bool) {
	b.checkBounds(index)

	bIndex, lIndex, mask := getIndexMask(index)
	b.locks[lIndex].Lock()

	if !value {
		b.Arr[bIndex] &= ^mask // Example: ANDed with 11111011
	} else {
		b.Arr[bIndex] |= mask // Example: ORed with 00010000
	}

	b.locks[lIndex].Unlock()
}

func (b BitSlice) FlipValue(index uint) {
	b.checkBounds(index)

	bIndex, lIndex, mask := getIndexMask(index)
	b.locks[lIndex].Lock()

	b.Arr[bIndex] ^= mask

	b.locks[lIndex].Unlock()
}
