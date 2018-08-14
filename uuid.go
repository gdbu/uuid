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
	// Make a byteslice with a capacity of stringLen
	// Note: The string length is greater than the hex length, this provides us space
	// to add our UUID separators after encoding to hexidecimal
	bs := make([]byte, stringLen)
	// Encode our UUID bytes as hexidecimal to our byteslice
	hex.Encode(bs, u[:])
	for i := 0; i < len(bs); i++ {
		// Switch on index
		switch i {
		// Separator index case
		case separator1, separator2, separator3, separator4:
			// Our index matches a separator index, prepend separator byte to this index
			prependByte(bs[i:], separator)
		}
	}

	// Convert byteslice to outbound string
	out = string(bs)
	return
}

// Time will return an associated time
func (u *UUID) Time() (out time.Time) {
	// Make a byteslice with a size of 8 bytes (same size as int64)
	bs := make([]byte, 8)
	// Copy first 8 bytes of UUID to our byteslice
	copy(bs, u[:8])
	// Convert the byteslice to an int64 and parse that value as a Unix timestamp
	return time.Unix(0, toValue(bs))
}

// MarshalJSON is a JSON encoding helper func
func (u UUID) MarshalJSON() (bs []byte, err error) {
	// Make the byteslice big enough for our string plus our quotation marks
	bs = make([]byte, 0, stringLen+2)
	// Append first quotation
	bs = append(bs, '"')
	// Append string representation of our UUID
	bs = append(bs, u.String()...)
	// Append second quotation
	bs = append(bs, '"')
	return
}

// UnmarshalJSON is a JSON decoding helper func
func (u *UUID) UnmarshalJSON(bs []byte) (err error) {
	var str string
	// Unmarshal inbound bytes as a string
	if err = json.Unmarshal(bs, &str); err != nil {
		return
	}

	// Convert parsed string to a byteslice
	strbs := []byte(str)
	// Remove separators from the string byteslice
	removeSeparators(strbs)
	// Decode cleaned bytes as Hexidecimal
	// Note: We can ignore any bytes which extend past our hexLen
	hex.Decode(u[:], strbs[:hexLen])
	return
}

func removeSeparators(bs []byte) {
	// Remove the four separators from the provided byteslice
	// Note: This will leave zero'd bytes for the removed indexes at the end of the byteslice.
	// It should be noted that we remove the separators in reverse-order so we ensure
	// we don't lose any crucial bytes.
	removeByte(bs, separator4)
	removeByte(bs, separator3)
	removeByte(bs, separator2)
	removeByte(bs, separator1)
}
