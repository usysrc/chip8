package emulator

import "fmt"

type Emulator struct {
	Memory *Memory
}

func NewEmulator() *Emulator {
	return &Emulator{
		Memory: &Memory{},
	}
}

func (e *Emulator) Run() {
	fmt.Println("Starting Chip8 Emulator")
}
