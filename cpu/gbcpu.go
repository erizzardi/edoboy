package cpu

// Type representing the 16 bit CPU registers.
// It contains the two individual 8-bit registers that compose the 16-bit one.
// More info on registers:
// https://gbdev.io/pandocs/CPU_Registers_and_Flags.html
type GBRegister struct {
	high, low byte
}

func (r *GBRegister) SetHigh(value byte) {
	r.high = value
}

func (r *GBRegister) SetLow(value byte) {
	r.low = value
}

func (r *GBRegister) Set(value uint16) {
	r.high = byte(value >> 8)
	r.low = byte(value)
}

func (r *GBRegister) GetHigh() byte {
	return r.high
}

func (r *GBRegister) GetLow() byte {
	return r.low
}

func (r *GBRegister) Get() uint16 {
	return uint16(r.high)<<8 | uint16(r.low)
}

// Type that implements the CPU interface.
// It holds all the registers.
type GBCpu struct {
	AF, BC, DE, HL, SP, PC *GBRegister
}

// This returns a new CPU, with the registers initialized
// to their default value, according to the documentation
func NewCpu() *GBCpu {
	return &GBCpu{
		AF: &GBRegister{high: 0x01, low: 0xB0},
		BC: &GBRegister{high: 0x00, low: 0x13},
		DE: &GBRegister{high: 0x00, low: 0xD8},
		HL: &GBRegister{high: 0x01, low: 0x4D},
		SP: &GBRegister{high: 0xFF, low: 0xFE},
		PC: &GBRegister{high: 0x01, low: 0x00},
	}
}
