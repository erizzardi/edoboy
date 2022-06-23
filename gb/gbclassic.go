package gb

import (
	"githun.com/erizzardi/edoboy/cpu"
	"githun.com/erizzardi/edoboy/ram"
)

// Naming sucks
// This implements the Gameboy interface
type GBClassic struct {
	Cpu cpu.Cpu
	Ram ram.Ram
}

func NewGB() Gameboy {
	return &GBClassic{
		Cpu: cpu.NewCpu(),
		Ram: ram.NewGBRam(),
	}
}

func (gb *GBClassic) Fetch()   {}
func (gb *GBClassic) Decode()  {}
func (gb *GBClassic) Execute() {}
