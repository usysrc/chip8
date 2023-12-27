package main

import (
	"log"

	"github.com/usysrc/chip8/emulator"
)

func main() {
	emulator := emulator.NewEmulator()
	err := emulator.LoadROM()
	if err != nil {
		log.Fatal(err)
	}
	err = emulator.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer emulator.Cleanup()
	err = emulator.Run()
	if err != nil {
		log.Fatal(err)
	}
}
