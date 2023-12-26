package main

import "github.com/usysrc/chip8/emulator"

func main() {
	emulator := emulator.NewEmulator()
	emulator.LoadROM()
	emulator.Run()
}
