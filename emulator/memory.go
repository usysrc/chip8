package emulator

import (
	"fmt"
	"log"
)

const (
	MemorySize   = 4096
	ProgramStart = 0x200
)

type Memory struct {
	content [MemorySize]byte
}

// read a byte from memory
func (m *Memory) Read(address uint16) (byte, error) {
	if address > MemorySize {
		return 0x00, fmt.Errorf("address out of bounds")
	}
	return m.content[address], nil
}

// write a byte to memory
func (m *Memory) Write(address uint16, value byte) error {
	if address > MemorySize {
		return fmt.Errorf("address out of bounds")
	}
	m.content[address] = value
	return nil
}

// fetch an opcode from memory
func (m *Memory) Opcode(address uint16) (uint16, error) {
	if address+1 > MemorySize {
		return 0, fmt.Errorf("address out of bounds")
	}

	fmt.Println(m.content[address], address)
	// the opcode is in stored in two bytes, so we need to get both
	high := uint16(m.content[address])
	low := uint16(m.content[address+1])
	// compose the bytes into highlow, the high byte being shifting left by 8 bit
	return high<<8 | low, nil
}

// load a segment to memory
func (m *Memory) Load(data []byte) error {
	if len(data) > MemorySize-ProgramStart {
		return fmt.Errorf("size exceeds available memory")
	}
	delta := copy(m.content[ProgramStart:], data)
	if delta == 0 {
		return fmt.Errorf("nothing loaded into memory")
	}
	log.Printf("wrote %d bytes to memory", delta)
	return nil
}
