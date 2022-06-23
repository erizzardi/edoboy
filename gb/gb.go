package gb

// This interface defines the Gameboy type.
// Implementations of this interface correspond to different gb architectures
// e.g. gameboy classic, color, advance etc.
// It's designed to wrap all the GB components up (e.g. cpu, ram etc.) and execute the high-level operations
// (i.e. fetch, decode, execute)
type Gameboy interface {
	// The Gameboy loops over three operations:
	// 	- Fetch: fetches the operation from memory
	// 	- Decode: decodes the operation (opcode) and returns the corresponding function
	// 	- Execute: executes the operation decoded in the preceeding step
	Fetch() byte
	Decode()
	Execute()
}
