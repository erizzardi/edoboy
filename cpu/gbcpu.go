package cpu

import "githun.com/erizzardi/edoboy/util"

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
	return util.Join(r.high, r.low)
}

// Type that implements the CPU interface.
// It holds all the registers.
type GBCpu struct {
	AF, BC, DE, HL, SP, PC *GBRegister
	z, n, hc, c            bool
}

func (cpu *GBCpu) SetZ(val bool) {
	cpu.z = val
}

func (cpu *GBCpu) SetN(val bool) {
	cpu.n = val
}

func (cpu *GBCpu) SetHC(val bool) {
	cpu.hc = val
}

func (cpu *GBCpu) GetHC() byte {
	return util.Btoi(cpu.hc)
}

func (cpu *GBCpu) SetC(val bool) {
	cpu.c = val
}

func (cpu *GBCpu) GetC() byte {
	return util.Btoi(cpu.c)
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

// 8-bit addition
func (cpu *GBCpu) Add8Bit(a, b byte) byte {
	sum := uint16(a + b)
	// gb.Cpu.AF.SetHigh(uint8(sum))
	innerSum := a&0x0f + b&0x0f
	cpu.SetZ(sum == 0)
	cpu.SetN(false)
	cpu.SetHC(innerSum&0x10 == 0x10)
	cpu.SetC(uint16(innerSum)&0x100 == 0x100)
	return byte(sum)
}

// 8 bit subtraction
func (cpu *GBCpu) Sub8bit(a, b byte) byte {
	diff := uint16(a) - uint16(b)
	innerDiff := a&0x0f - b&0x0f
	cpu.SetZ(diff == 0)
	cpu.SetN(true)
	cpu.SetHC(innerDiff&0x10 == 0x10)
	cpu.SetC(int16(a)-int16(b) < 0)
	return byte(diff)
}

// 16-bit addition
