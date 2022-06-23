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

func (gb *GBClassic) Decode()  {}
func (gb *GBClassic) Execute() {}
