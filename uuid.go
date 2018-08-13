package uuid

import (
	"encoding/hex"
	"encoding/json"
	"time"
)

const (
	byteLen   = 14
	stringLen = 32
	separator = '-'

	separator1 = 8
	separator2 = 13
	separator3 = 18
	separator4 = 23
)

var (
	hexLen = hex.EncodedLen(byteLen)
	global = NewGenerator()
)

// New will return a new UUID utilizing the global generator
func New() UUID {
	return global.New()
}

// newUUID will return a new UUID from a provided value
// Note: The provided value will be truncated from 8 bytes to 6 bytes
func newUUID(value int64) (u UUID) {
	ts := toBytes(time.Now().UnixNano())
	copy(u[:8], ts)
	copy(u[8:], toBytes(value)[:6])
	return
}

// UUID represents a unique user identifier
type UUID [byteLen]byte

// String will return the string representation of a UUID
func (u *UUID) String() (out string) {
	bs := make([]byte, stringLen)
	hex.Encode(bs, u[:])
	for i := 0; i < len(bs); i++ {
		switch i {
		case separator1, separator2, separator3, separator4:
			prependByte(bs[i:], separator)
		}
	}

	out = string(bs)
	return
}

// Time will return an associated time
func (u *UUID) Time() (out time.Time) {
	bs := make([]byte, 8)
	copy(bs, u[:8])
	return time.Unix(0, toValue(bs))
}

// MarshalJSON is a JSON encoding helper func
func (u UUID) MarshalJSON() (bs []byte, err error) {
	bs = make([]byte, 0, stringLen+2)
	bs = append(bs, '"')
	bs = append(bs, u.String()...)
	bs = append(bs, '"')
	return
}

// UnmarshalJSON is a JSON decoding helper func
func (u *UUID) UnmarshalJSON(bs []byte) (err error) {
	var str string
	if err = json.Unmarshal(bs, &str); err != nil {
		return
	}

	strbs := []byte(str)
	removeByte(strbs, separator4)
	removeByte(strbs, separator3)
	removeByte(strbs, separator2)
	removeByte(strbs, separator1)
	hex.Decode(u[:], strbs)
	return
}
