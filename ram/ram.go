package ram

// Interface that defines the memory addresses.
// Implementations correspond to different GB models.
type Ram interface {
	//
	//
	// Write writes 'value' into 'address'
	Write(address uint16, value byte)
	//
	//
	// Read returns the memory value stored at 'address'
	Read(address uint16) byte
}
