package util

// Join joins two 8-bit numbers in one 16-bit value
func Join(high, low byte) uint16 {
	return uint16(high)<<8 | uint16(low)
}
