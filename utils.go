package uuid

import (
	crand "crypto/rand"
	"encoding/binary"
	"unsafe"
)

func toBytes(value int64) (bs []byte) {
	return (*[8]byte)(unsafe.Pointer(&value))[:]
}

func toValue(bs []byte) (value int64) {
	return *(*int64)(unsafe.Pointer(&bs[0]))
}

func prependByte(bs []byte, b byte) {
	for i := len(bs) - 2; i > -1; i-- {
		bs[i+1] = bs[i]
	}

	bs[0] = b
}

func removeByte(bs []byte, i int) {
	if i < 0 {
		// This should NEVER happen, if it happens - it's worth a panic
		panic("invalid index, cannot be negative")
	}

	for ; i < len(bs)-1; i++ {
		bs[i] = bs[i+1]
	}

	bs[len(bs)-1] = 0
}

func generateSeed() (seed int64) {
	var buf [8]byte
	crand.Read(buf[:])
	return int64(binary.LittleEndian.Uint64(buf[:]))
}
