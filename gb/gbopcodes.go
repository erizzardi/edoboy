package gb

import (
	"githun.com/erizzardi/edoboy/util"
)

// Opcodes for the original gameboy and gameboy color.
// A full list can be found here:
// http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf page 65

// For ease of access, every operation is written as a function of the gameboy
// (conceptually doesn't make much sense...) but allows to cluster them together in a
// map[byte]func structure, so the Decode() function contains only a map lookup
var operations = map[byte]func(*GBClassic){

	//=====================================================================================================
	// Caveats on functions
	// ----------------------------------------------------------------------------------------------------
	// gb.Fetch() is to be used to fetch the immediate next value in memory - it automatically increases PC
	// gb.Ram.Read() is to read an arbitrary address in memory. It doesn't require a PC increase
	//=====================================================================================================

	//============
	// 8-bit loads
	//============
	// LD B,d8
	0x06: func(gb *GBClassic) {
		value := gb.Fetch()
		gb.Cpu.BC.SetHigh(value)
	},
	// LD C,d8
	0x0E: func(gb *GBClassic) {
		value := gb.Fetch()
		gb.Cpu.BC.SetLow(value)
	},
	// LD D,d8
	0x16: func(gb *GBClassic) {
		value := gb.Fetch()
		gb.Cpu.DE.SetHigh(value)
	},
	// LD E,d8
	0x1E: func(gb *GBClassic) {
		value := gb.Fetch()
		gb.Cpu.DE.SetLow(value)
	},
	// LD D,d8
	0x26: func(gb *GBClassic) {
		value := gb.Fetch()
		gb.Cpu.HL.SetHigh(value)
	},
	// LD E,d8
	0x2E: func(gb *GBClassic) {
		value := gb.Fetch()
		gb.Cpu.HL.SetLow(value)
	},
	// LD A,A
	0x7F: func(gb *GBClassic) {
		gb.Cpu.AF.SetHigh(gb.Cpu.AF.GetHigh())
	},
	// LD A,B
	0x78: func(gb *GBClassic) {
		gb.Cpu.AF.SetHigh(gb.Cpu.BC.GetHigh())
	},
	// LD A,C
	0x79: func(gb *GBClassic) {
		gb.Cpu.AF.SetHigh(gb.Cpu.BC.GetLow())
	},
	// LD A,D
	0x7A: func(gb *GBClassic) {
		gb.Cpu.AF.SetHigh(gb.Cpu.DE.GetHigh())
	},
	// LD A,E
	0x7B: func(gb *GBClassic) {
		gb.Cpu.AF.SetHigh(gb.Cpu.DE.GetLow())
	},
	// LD A,H
	0x7C: func(gb *GBClassic) {
		gb.Cpu.AF.SetHigh(gb.Cpu.HL.GetHigh())
	},
	// LD A,L
	0x7D: func(gb *GBClassic) {
		gb.Cpu.AF.SetHigh(gb.Cpu.HL.GetLow())
	},
	// LD A,(HL)
	0x7E: func(gb *GBClassic) {
		gb.Cpu.AF.SetHigh(gb.Ram.Read(gb.Cpu.HL.Get()))
	},
	// LD B,B
	0x40: func(gb *GBClassic) {
		gb.Cpu.BC.SetHigh(gb.Cpu.BC.GetHigh())
	},
	// LD B,C
	0x41: func(gb *GBClassic) {
		gb.Cpu.BC.SetHigh(gb.Cpu.BC.GetLow())
	},
	// LD B,D
	0x42: func(gb *GBClassic) {
		gb.Cpu.BC.SetHigh(gb.Cpu.DE.GetHigh())
	},
	// LD B,E
	0x43: func(gb *GBClassic) {
		gb.Cpu.BC.SetHigh(gb.Cpu.DE.GetLow())
	},
	// LD B,H
	0x44: func(gb *GBClassic) {
		gb.Cpu.BC.SetHigh(gb.Cpu.HL.GetHigh())
	},
	// LD B,L
	0x45: func(gb *GBClassic) {
		gb.Cpu.BC.SetHigh(gb.Cpu.HL.GetLow())
	},
	// LD B,(HL)
	0x46: func(gb *GBClassic) {
		gb.Cpu.BC.SetHigh(gb.Ram.Read(gb.Cpu.HL.Get()))
	},
	// LD C,B
	0x48: func(gb *GBClassic) {
		gb.Cpu.BC.SetLow(gb.Cpu.BC.GetHigh())
	},
	// LD C,C
	0x49: func(gb *GBClassic) {
		gb.Cpu.BC.SetLow(gb.Cpu.BC.GetLow())
	},
	// LD C,D
	0x4A: func(gb *GBClassic) {
		gb.Cpu.BC.SetLow(gb.Cpu.DE.GetHigh())
	},
	// LD C,E
	0x4B: func(gb *GBClassic) {
		gb.Cpu.BC.SetLow(gb.Cpu.DE.GetLow())
	},
	// LD C,H
	0x4C: func(gb *GBClassic) {
		gb.Cpu.BC.SetLow(gb.Cpu.HL.GetHigh())
	},
	// LD C,L
	0x4D: func(gb *GBClassic) {
		gb.Cpu.BC.SetLow(gb.Cpu.HL.GetLow())
	},
	// LD C,(HL)
	0x4E: func(gb *GBClassic) {
		gb.Cpu.BC.SetLow(gb.Ram.Read(gb.Cpu.HL.Get()))
	},
	// LD D,B
	0x50: func(gb *GBClassic) {
		gb.Cpu.DE.SetHigh(gb.Cpu.BC.GetHigh())
	},
	// LD D,C
	0x51: func(gb *GBClassic) {
		gb.Cpu.DE.SetHigh(gb.Cpu.BC.GetLow())
	},
	// LD D,D
	0x52: func(gb *GBClassic) {
		gb.Cpu.DE.SetHigh(gb.Cpu.DE.GetHigh())
	},
	// LD D,E
	0x53: func(gb *GBClassic) {
		gb.Cpu.DE.SetHigh(gb.Cpu.DE.GetLow())
	},
	// LD D,H
	0x54: func(gb *GBClassic) {
		gb.Cpu.DE.SetHigh(gb.Cpu.HL.GetHigh())
	},
	// LD D,L
	0x55: func(gb *GBClassic) {
		gb.Cpu.DE.SetHigh(gb.Cpu.HL.GetLow())
	},
	// LD D,(HL)
	0x56: func(gb *GBClassic) {
		gb.Cpu.DE.SetHigh(gb.Ram.Read(gb.Cpu.HL.Get()))
	},
	// LD E,B
	0x58: func(gb *GBClassic) {
		gb.Cpu.DE.SetLow(gb.Cpu.BC.GetHigh())
	},
	// LD E,C
	0x59: func(gb *GBClassic) {
		gb.Cpu.DE.SetLow(gb.Cpu.BC.GetLow())
	},
	// LD E,D
	0x5A: func(gb *GBClassic) {
		gb.Cpu.DE.SetLow(gb.Cpu.DE.GetHigh())
	},
	// LD E,E
	0x5B: func(gb *GBClassic) {
		gb.Cpu.DE.SetLow(gb.Cpu.DE.GetLow())
	},
	// LD E,H
	0x5C: func(gb *GBClassic) {
		gb.Cpu.DE.SetLow(gb.Cpu.HL.GetHigh())
	},
	// LD E,L
	0x5D: func(gb *GBClassic) {
		gb.Cpu.DE.SetLow(gb.Cpu.HL.GetLow())
	},
	// LD E,(HL)
	0x5E: func(gb *GBClassic) {
		gb.Cpu.DE.SetLow(gb.Ram.Read(gb.Cpu.HL.Get()))
	},
	// LD H,B
	0x60: func(gb *GBClassic) {
		gb.Cpu.HL.SetHigh(gb.Cpu.BC.GetHigh())
	},
	// LD H,C
	0x61: func(gb *GBClassic) {
		gb.Cpu.HL.SetHigh(gb.Cpu.BC.GetLow())
	},
	// LD H,D
	0x62: func(gb *GBClassic) {
		gb.Cpu.HL.SetHigh(gb.Cpu.DE.GetHigh())
	},
	// LD H,E
	0x63: func(gb *GBClassic) {
		gb.Cpu.HL.SetHigh(gb.Cpu.DE.GetLow())
	},
	// LD H,H
	0x64: func(gb *GBClassic) {
		gb.Cpu.HL.SetHigh(gb.Cpu.HL.GetHigh())
	},
	// LD H,L
	0x65: func(gb *GBClassic) {
		gb.Cpu.HL.SetHigh(gb.Cpu.HL.GetLow())
	},
	// LD H,(HL)
	0x66: func(gb *GBClassic) {
		gb.Cpu.HL.SetHigh(gb.Ram.Read(gb.Cpu.HL.Get()))
	},
	// LD L,B
	0x68: func(gb *GBClassic) {
		gb.Cpu.HL.SetLow(gb.Cpu.BC.GetHigh())
	},
	// LD L,C
	0x69: func(gb *GBClassic) {
		gb.Cpu.HL.SetLow(gb.Cpu.BC.GetLow())
	},
	// LD L,D
	0x6A: func(gb *GBClassic) {
		gb.Cpu.HL.SetLow(gb.Cpu.DE.GetHigh())
	},
	// LD L,E
	0x6B: func(gb *GBClassic) {
		gb.Cpu.HL.SetLow(gb.Cpu.DE.GetLow())
	},
	// LD L,H
	0x6C: func(gb *GBClassic) {
		gb.Cpu.HL.SetLow(gb.Cpu.HL.GetHigh())
	},
	// LD L,L
	0x6D: func(gb *GBClassic) {
		gb.Cpu.HL.SetLow(gb.Cpu.HL.GetLow())
	},
	// LD L,(HL)
	0x6E: func(gb *GBClassic) {
		gb.Cpu.HL.SetLow(gb.Ram.Read(gb.Cpu.HL.Get()))
	},
	// LD (HL),B
	0x70: func(gb *GBClassic) {
		gb.Ram.Write(gb.Cpu.HL.Get(), gb.Cpu.BC.GetHigh())
	},
	// LD (HL),C
	0x71: func(gb *GBClassic) {
		gb.Ram.Write(gb.Cpu.HL.Get(), gb.Cpu.BC.GetLow())
	},
	// LD (HL),D
	0x72: func(gb *GBClassic) {
		gb.Ram.Write(gb.Cpu.HL.Get(), gb.Cpu.DE.GetHigh())
	},
	// LD (HL),E
	0x73: func(gb *GBClassic) {
		gb.Ram.Write(gb.Cpu.HL.Get(), gb.Cpu.DE.GetLow())
	},
	// LD (HL),H
	0x74: func(gb *GBClassic) {
		gb.Ram.Write(gb.Cpu.HL.Get(), gb.Cpu.HL.GetHigh())
	},
	// LD (HL),L
	0x75: func(gb *GBClassic) {
		gb.Ram.Write(gb.Cpu.HL.Get(), gb.Cpu.HL.GetLow())
	},
	// LD (HL),d8
	0x36: func(gb *GBClassic) {
		gb.Ram.Write(gb.Cpu.HL.Get(), gb.Fetch())
	},
	// LD A,(BC)
	0x0A: func(gb *GBClassic) {
		gb.Cpu.AF.SetHigh(gb.Ram.Read(gb.Cpu.BC.Get()))
	},
	// LD A,(DE)
	0x1A: func(gb *GBClassic) {
		gb.Cpu.AF.SetHigh(gb.Ram.Read(gb.Cpu.DE.Get()))
	},
	// LD A,(a16)
	0xFA: func(gb *GBClassic) {
		low := gb.Fetch()
		high := gb.Fetch()
		gb.Cpu.AF.SetHigh(gb.Ram.Read(util.Join(high, low)))
	},
	// LD A,d8
	0x3E: func(gb *GBClassic) {
		gb.Cpu.AF.SetHigh(gb.Fetch())
	},
	// LD B,A
	0x47: func(gb *GBClassic) {
		gb.Cpu.BC.SetHigh(gb.Cpu.AF.GetHigh())
	},
	// LD C,A
	0x4F: func(gb *GBClassic) {
		gb.Cpu.BC.SetLow(gb.Cpu.AF.GetHigh())
	},
	// LD D,A
	0x57: func(gb *GBClassic) {
		gb.Cpu.DE.SetHigh(gb.Cpu.AF.GetHigh())
	},
	// LD E,A
	0x5F: func(gb *GBClassic) {
		gb.Cpu.DE.SetLow(gb.Cpu.AF.GetHigh())
	},
	// LD H,A
	0x67: func(gb *GBClassic) {
		gb.Cpu.HL.SetHigh(gb.Cpu.AF.GetHigh())
	},
	// LD L,A
	0x6F: func(gb *GBClassic) {
		gb.Cpu.HL.SetLow(gb.Cpu.AF.GetHigh())
	},
	// LD (BC),A
	0x02: func(gb *GBClassic) {
		gb.Ram.Write(gb.Cpu.BC.Get(), gb.Cpu.AF.GetHigh())
	},
	// LD (DE),A
	0x12: func(gb *GBClassic) {
		gb.Ram.Write(gb.Cpu.DE.Get(), gb.Cpu.AF.GetHigh())
	},
	// LD (HL),A
	0x77: func(gb *GBClassic) {
		gb.Ram.Write(gb.Cpu.HL.Get(), gb.Cpu.AF.GetHigh())
	},
	// LD (a16),A
	0xEA: func(gb *GBClassic) {
		low := gb.Fetch()
		high := gb.Fetch()
		gb.Ram.Write(util.Join(high, low), gb.Cpu.AF.GetHigh())
	},
	// LD A,(C)
	0xF2: func(gb *GBClassic) {
		val := gb.Ram.Read(0xFF00 + uint16(gb.Cpu.BC.GetLow()))
		gb.Cpu.AF.SetHigh(val)
	},
	// LD (C),A
	0xE2: func(gb *GBClassic) {
		gb.Ram.Write(0xFF00+uint16(gb.Cpu.BC.GetLow()), gb.Cpu.AF.GetHigh())
	},
	// LDD A,(HL)
	0x3A: func(gb *GBClassic) {
		gb.Cpu.AF.SetHigh(gb.Ram.Read(gb.Cpu.HL.Get()))
		gb.Cpu.HL.Set(gb.Cpu.HL.Get() - 1)
	},
	// LDD (HL),A
	0x32: func(gb *GBClassic) {
		gb.Ram.Write(gb.Cpu.HL.Get(), gb.Cpu.AF.GetHigh())
		gb.Cpu.HL.Set(gb.Cpu.HL.Get() - 1)
	},
	// LDI A,(HL)
	0x2A: func(gb *GBClassic) {
		gb.Cpu.AF.SetHigh(gb.Ram.Read(gb.Cpu.HL.Get()))
		gb.Cpu.HL.Set(gb.Cpu.HL.Get() + 1)
	},
	// LDI (HL),A
	0x22: func(gb *GBClassic) {
		gb.Ram.Write(gb.Cpu.HL.Get(), gb.Cpu.AF.GetHigh())
		gb.Cpu.HL.Set(gb.Cpu.HL.Get() + 1)
	},
	// LDH (a8),A
	0xE0: func(gb *GBClassic) {
		val := gb.Fetch()
		gb.Ram.Write(0xFF00+uint16(val), gb.Cpu.AF.GetHigh())
	},
	// LDH A,(a8)
	0xF0: func(gb *GBClassic) {
		val := gb.Fetch()
		gb.Cpu.AF.SetHigh(gb.Ram.Read(0xFF00 + uint16(val)))
	},

	//=============
	// 16-bit loads
	//=============
	//LD BC,d16
	0x01: func(gb *GBClassic) {
		low := gb.Fetch()
		high := gb.Fetch()
		gb.Cpu.BC.Set(util.Join(high, low))
	},
	// LD DE,d16
	0x11: func(gb *GBClassic) {
		low := gb.Fetch()
		high := gb.Fetch()
		gb.Cpu.DE.Set(util.Join(high, low))
	},
	// LD HL,16
	0x21: func(gb *GBClassic) {
		low := gb.Fetch()
		high := gb.Fetch()
		gb.Cpu.HL.Set(util.Join(high, low))
	},
	// LD SP,d16
	0x31: func(gb *GBClassic) {
		low := gb.Fetch()
		high := gb.Fetch()
		gb.Cpu.SP.Set(util.Join(high, low))
	},
	// LD SP,HL
	0xF9: func(gb *GBClassic) {
		gb.Cpu.SP.Set(gb.Cpu.HL.Get())
	},
	// LDHL SP,s8
	0xF8: func(gb *GBClassic) {
		// this operation reads the number as signed
		val := int16(gb.Fetch())
		sum := uint16(val + int16(gb.Cpu.SP.Get()))
		gb.Cpu.HL.Set(sum)
		// check for C and HC flags
		innerSum := (val & 0x0F) + (int16(gb.Cpu.SP.Get()) & 0x0F)
		gb.Cpu.SetZ(false)
		gb.Cpu.SetN(false)
		gb.Cpu.SetHC((innerSum & 0x10) == 0x10)
		gb.Cpu.SetC((innerSum & 0x100) == 0x100)

	},
	// LD (a16),SP
	0x08: func(gb *GBClassic) {
		low := gb.Fetch()
		high := gb.Fetch()
		gb.Ram.Write(util.Join(high, low), gb.Cpu.SP.GetLow())
		gb.Ram.Write(util.Join(high, low)+1, gb.Cpu.SP.GetHigh())
	},
	// PUSH AF
	0xF5: func(gb *GBClassic) {
		gb.PushStack(gb.Cpu.AF.GetHigh())
		gb.PushStack(gb.Cpu.AF.GetLow())
	},
	// PUSH BC
	0xC5: func(gb *GBClassic) {
		gb.PushStack(gb.Cpu.BC.GetHigh())
		gb.PushStack(gb.Cpu.BC.GetLow())
	},
	// PUSH DE
	0xD5: func(gb *GBClassic) {
		gb.PushStack(gb.Cpu.DE.GetHigh())
		gb.PushStack(gb.Cpu.DE.GetLow())
	},
	// PUSH HL
	0xE5: func(gb *GBClassic) {
		gb.PushStack(gb.Cpu.HL.GetHigh())
		gb.PushStack(gb.Cpu.HL.GetLow())
	},
	// POP AF
	0xF1: func(gb *GBClassic) {
		gb.Cpu.AF.SetLow(gb.PopStack())
		gb.Cpu.AF.SetHigh(gb.PopStack())
	},
	// POP BC
	0xC1: func(gb *GBClassic) {
		gb.Cpu.BC.SetLow(gb.PopStack())
		gb.Cpu.BC.SetHigh(gb.PopStack())
	},
	// POP DE
	0xD1: func(gb *GBClassic) {
		gb.Cpu.DE.SetLow(gb.PopStack())
		gb.Cpu.DE.SetHigh(gb.PopStack())
	},
	// POP HL
	0xE1: func(gb *GBClassic) {
		gb.Cpu.HL.SetLow(gb.PopStack())
		gb.Cpu.HL.SetHigh(gb.PopStack())
	},

	//==========
	// 8-bit ALU
	//==========
	// ADD A,A
}
