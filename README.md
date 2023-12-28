# CHIP-8 Emulator

A simple CHIP-8 emulator written in Go.

## Overview

This project is a CHIP-8 emulator implemented in the Go programming language. CHIP-8 is an interpreted programming language designed in the 1970s. CHIP-8 programs are run on a virtual machine, and this emulator simulates that virtual machine.

## Features

- Emulation of the CHIP-8 virtual machine.
- Basic CPU, memory, and display functionality.
- Support for executing CHIP-8 ROMs.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/usysrc/chip8
   ```

2. Change into the project directory:

   ```bash
   cd chip8
   ```

3. Build and run the emulator:

   ```bash
   go run main.go 
   ```

## Usage

- Replace `path/to/your/rom.ch8` with the path to the CHIP-8 ROM you want to run.
- Use the keyboard keys mapped to the CHIP-8 keypad for input.

## Key Mapping

The default key mapping for the CHIP-8 keypad is as follows:

```
1 2 3 C
4 5 6 D
7 8 9 E
A 0 B F
```

## Credits

- This emulator is based on the CHIP-8 technical reference and specifications.
- Special thanks to the CHIP-8 community for preserving and sharing information about this classic system.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

