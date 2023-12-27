package emulator

import (
	"fmt"
	"os"
)

type Emulator struct {
	Memory  *Memory
	Display *Display
}

func NewEmulator() *Emulator {
	return &Emulator{
		Memory:  &Memory{},
		Display: &Display{},
	}
}

func (e *Emulator) Init() error {
	err := e.Display.Init()
	return err
}

func (e *Emulator) Cleanup() {
	e.Display.Close()
}

func (e *Emulator) Run() error {
	fmt.Println("Starting Chip8 Emulator")
	for {
		err := e.Display.Draw(32, 32)
		if err != nil {
			return err
		}
	}
}

func (e *Emulator) LoadROM() error {
	file, err := os.ReadFile("roms/INVADERS")
	if err != nil {
		return err
	}
	err = e.Memory.Load(file)
	if err != nil {
		return err
	}
	return nil
}
