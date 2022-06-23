package cpu

// Type representing the CPU registers.
// I decided to use three distinct values for the 16-bit full register
// and the two 8-bit physical components. It's easier to write, less error-prone
// and it costs only 16 bits of memory per register (96 bits in total).
// More info on registers:
// https://gbdev.io/pandocs/CPU_Registers_and_Flags.html
type GBRegister struct {
	sixb      uint16
	high, low byte
}

func (r *GBRegister) SetHigh(value byte) {
	r.high = value
}

func (r *GBRegister) SetLow(value byte) {
	r.low = value
}

func (r *GBRegister) Set(value uint16) {
	r.sixb = value
}

func (r *GBRegister) GetHigh() byte {
	return r.high
}

func (r *GBRegister) GetLow() byte {
	return r.low
}

func (r *GBRegister) Get() uint16 {
	return r.sixb
}

// Type that implements the CPU interface.
// It holds all the registers.
type GBCpu struct {
	AF, BC, DE, HL, SP, PC *GBRegister
}

func NewCpu() Cpu {
	return &GBCpu{
		AF: &GBRegister{},
		BC: &GBRegister{},
		DE: &GBRegister{},
		HL: &GBRegister{},
		SP: &GBRegister{},
		PC: &GBRegister{},
	}
}
