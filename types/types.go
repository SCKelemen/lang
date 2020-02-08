package types

// SanityCheck produces checked + message given message
func SanityCheck(message string) string {
	return "checked " + message
}

type Primitive int

const (
	Invalid Primitive = iota

	None
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64

	Byte = Uint8
	Char = Uint8
	Rune = Int32
)
