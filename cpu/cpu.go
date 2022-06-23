package cpu

// Interface representing the GB CPU.
// Implementation of this interface correspond to different cpu architectures
// e.g. LR35902 (classic, color),  ARM7TDMI (advance)
type Cpu interface{}

// // Interface representing a GB CPU register
// type Register interface {
// 	SetHigh(value byte)
// 	SetLow(value byte)
// 	Set(value uint16)
// }
