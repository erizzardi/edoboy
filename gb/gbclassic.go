package gb

import (
	"githun.com/erizzardi/edoboy/cpu"
	"githun.com/erizzardi/edoboy/ram"
)

// Naming sucks
// This implements the Gameboy interface
type GBClassic struct {
	Cpu *cpu.GBCpu
	Ram *ram.GBRam

	op      byte
	op16bit byte
}

func NewGBClassic() Gameboy {
	return &GBClassic{
		Cpu: cpu.NewCpu(),
		Ram: ram.NewGBRam(),
	}
}

// Fetch takes the 8-bit instruction located at memory address PC
// and stores it in gb.op.
func (gb *GBClassic) Fetch() byte {
	op := gb.Ram.Read(gb.Cpu.PC.Get())
	// advance PC registry by 1
	gb.Cpu.PC.Set(gb.Cpu.PC.Get() + 1)
	return op
}

func (gb *GBClassic) Decode() {

}
func (gb *GBClassic) Execute() {}

//============
// Miscellanea
//============

// Reads a value from the stack and
// decreases SP by 1
func (gb *GBClassic) PopStack() uint8 {
	ret := gb.Ram.Read(gb.Cpu.SP.Get() + 1)
	gb.Cpu.SP.Set(gb.Cpu.SP.Get() + 1)
	return ret
}

// Writes a value to the stack and
// increases SP by 1
func (gb *GBClassic) PushStack(val uint8) {
	gb.Ram.Write(gb.Cpu.SP.Get()+1, val)
	gb.Cpu.SP.Set(gb.Cpu.SP.Get() + 1)
}
