package ram

type GBRam struct {

	// This represents the GB memory array.
	// The GB has a 16 bit address bus (0000-FFFF), and 132KB total memory.
	// This array is used for ROM, RAM and I/O.
	// For info on memory mapping:
	// https://gbdev.io/pandocs/Memory_Map.html
	deference [0xFFFF]byte // C ftw
}

func NewGBRam() *GBRam {
	return &GBRam{}
}

func (ram *GBRam) Write(address uint16, value byte) {
	ram.deference[address] = value
}

func (ram *GBRam) Read(address uint16) byte {
	return ram.deference[address]
}
