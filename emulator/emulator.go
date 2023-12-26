package emulator

import (
	"fmt"
	"io/ioutil"
)

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

func (e *Emulator) LoadROM() error {
	file, err := ioutil.ReadFile("roms/INVADER")
	if err != nil {
		return err
	}
	err = e.Memory.Load(file)
	if err != nil {
		return err
	}
	return nil
}
