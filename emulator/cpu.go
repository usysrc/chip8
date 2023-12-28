package emulator

import (
	"fmt"
	"log"
	"math/rand"
)

// CPU represents the CHIP-8 CPU.
type CPU struct {
	// Registers
	vRegisters [16]byte // V0 to VF registers
	iRegister  uint16   // Index register
	pc         uint16   // Program counter

	// Timers
	delayTimer byte
	soundTimer byte

	// Stack
	stack [16]uint16
	sp    byte // Stack pointer
}

// NewCPU creates and initializes a new CHIP-8 CPU.
func NewCPU() *CPU {
	return &CPU{
		vRegisters: [16]byte{},
		iRegister:  0,
		pc:         ProgramStart,

		delayTimer: 0x00,
		soundTimer: 0x00,

		stack: [16]uint16{},
		sp:    0,
	}
}

// ExecuteCycle executes a single cycle of the CHIP-8 CPU.
func (cpu *CPU) ExecuteCycle(memory *Memory, display *Display) error {
	// Fetch opcode from memory
	opcode, err := memory.Opcode(cpu.pc)
	if err != nil {
		return err
	}

	// Decode and execute the opcode
	cpu.decodeAndExecute(opcode, memory, display)

	// Update timers if needed
	cpu.updateTimers()

	// increment pc
	cpu.pc++

	return nil
}

// decodeAndExecute decodes and executes a CHIP-8 opcode.
func (cpu *CPU) decodeAndExecute(opcode uint16, memory *Memory, display *Display) {
	// Extract parts of the opcode
	x := (opcode & 0x0F00) >> 8 // Extract register index x
	y := (opcode & 0x00F0) >> 8 // Extract register index y
	z := (opcode & 0x000F) >> 8 // Extract the subcommand
	kk := byte(opcode & 0x00FF) // Extract byte value
	n := int(opcode & 0x000F)   // Extract height of the sprite
	nnn := opcode & 0x0FFF      // Extract three nibble
	address := opcode & 0x0FFF  // Extract the address

	// Decode based on the opcode
	switch opcode & 0xF000 {
	case 0x1000:
		// Opcode 1nnn: Jump to address nnn
		cpu.pc = address
	case 0x2000:
		// Opcode 2nnn: Call subroutine at nnn
		// Increment stackpointer
		cpu.sp++

		// Save current program counter on the stack
		cpu.stack[cpu.sp] = cpu.pc

		// Set program counter to address nnn
		cpu.pc = address
	case 0x3000:
		if cpu.vRegisters[x] == kk {
			cpu.pc += 2
		}
	case 0x4000:
		if cpu.vRegisters[x] != kk {
			cpu.pc += 2
		}
	case 0x5000:
		if cpu.vRegisters[x] == cpu.vRegisters[y] {
			cpu.pc += 2
		}
	case 0x6000:
		cpu.vRegisters[x] = kk
	case 0x7000:
		cpu.vRegisters[x] += kk
	case 0x8000:
		if z == 0 {
			cpu.vRegisters[x] = cpu.vRegisters[y]
		}
		if z == 1 {
			cpu.vRegisters[x] |= cpu.vRegisters[y]
		}
		if z == 2 {
			cpu.vRegisters[x] &= cpu.vRegisters[y]
		}
		if z == 3 {
			cpu.vRegisters[x] ^= cpu.vRegisters[y]
		}
		if z == 4 {
			cpu.vRegisters[x] += cpu.vRegisters[y]
		}
		if z == 5 {
			cpu.vRegisters[x] -= cpu.vRegisters[y]
		}
		if z == 6 {
			cpu.vRegisters[0xF] = cpu.vRegisters[y] & 0x01 // Store LSB of Vy in VF
			cpu.vRegisters[x] = cpu.vRegisters[y] >> 1     // Shift Vy right, store result in Vx
		}
		if z == 7 {
			if cpu.vRegisters[y] > cpu.vRegisters[x] {
				cpu.vRegisters[0xF] = 0 // Borrow
			} else {
				cpu.vRegisters[0xF] = 1 // No borrow
			}

			cpu.vRegisters[x] = cpu.vRegisters[y] - cpu.vRegisters[x]
		}
		if z == 0xE {
			cpu.vRegisters[0xF] = (cpu.vRegisters[y] & 0x80) >> 7 // Store MSB of Vy in VF
			cpu.vRegisters[x] = cpu.vRegisters[y] << 1            // Shift Vy left, store result in Vx
		}
	case 0x9000:
		if cpu.vRegisters[x] != cpu.vRegisters[y] {
			cpu.pc += 2
		}
	case 0xA000:
		cpu.iRegister = nnn
	case 0xB000:
		cpu.pc = nnn + uint16(cpu.vRegisters[0])
	case 0xC000:
		cpu.vRegisters[x] = byte(rand.Intn(256)) & kk
	case 0xD000:
		// Opcode Dxyn: Display n-byte sprite starting at memory location I at (Vx, Vy)
		// Set VF = collision
		vx := int(cpu.vRegisters[x])
		vy := int(cpu.vRegisters[y])

		// Reset collision flag
		cpu.vRegisters[0xF] = 0
		for i := 0; i < n; i++ {
			// Read a byte from memory
			spriteByte, err := memory.Read(cpu.iRegister + uint16(i))
			if err != nil {
				// TODO: make function return error
				log.Fatal(err)
			}
			// Iterate over each pixel in the byte
			for j := 0; j < 8; j++ {
				px := (vx + j) % DisplayWidth
				py := (vy + i) % DisplayHeight
				// Check if the pixel is turned on
				pixelOn := (spriteByte & (0x80 >> j)) != 0
				// XOR the pixel with the display
				existingPixel, err := display.Get(uint16(px), uint16(py))
				if err != nil {
					// TODO: make this function return an error
					log.Fatal(err)
				}
				newPixel := existingPixel != pixelOn

				// If the pixel is turned off due to XOR, set collision flag
				if existingPixel && !newPixel {
					cpu.vRegisters[0xF] = 1
				}

				err = display.Set(uint16(px), uint16(px), newPixel)
				if err != nil {
					// TODO: make this function return an error
					log.Fatal(err)
				}
			}
		}
		// TODO: add missing opcodes
	default:
		// Unknown opcode, handle or log error
		fmt.Printf("Unknown opcode: 0x%X\n", opcode)
	}
}

// updateTimers updates the delay and sound timers.
func (cpu *CPU) updateTimers() {
	// Implement timer update logic
}
