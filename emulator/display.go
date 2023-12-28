package emulator

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	DotSize       = int32(8)
	DisplayWidth  = 64
	DisplayHeight = 32
)

type Display struct {
	pixels [DisplayWidth * DisplayHeight]bool
}

func (d *Display) Init() error {
	rl.InitWindow(800, 450, "chip8")
	rl.SetTargetFPS(60)
	return nil
}

func (d *Display) Clear() {
	for i := 0; i < DisplayWidth; i++ {
		for j := 0; j < DisplayHeight; j++ {
			d.pixels[j*DisplayWidth+i] = false
		}
	}
}

func (d *Display) Get(x, y uint16) (bool, error) {
	if x < 0 || x > DisplayWidth || y < 0 || y > DisplayHeight {
		return false, fmt.Errorf("trying to draw out of screen bounds")
	}
	return d.pixels[y*DisplayWidth+x], nil
}

func (d *Display) Set(x, y uint16, value bool) error {
	if x < 0 || x > DisplayWidth || y < 0 || y > DisplayHeight {
		return fmt.Errorf("trying to draw out of screen bounds")
	}
	d.pixels[y*DisplayWidth+x] = value
	return nil
}

func (d *Display) Draw() error {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	for i := 0; i < DisplayWidth; i++ {
		for j := 0; j < DisplayHeight; j++ {
			if d.pixels[j*DisplayWidth+i] {
				rl.DrawRectangle(int32(i)*DotSize, int32(j)*DotSize, DotSize, DotSize, rl.White)
			}
		}
	}

	rl.EndDrawing()
	return nil
}

func (d *Display) Close() {
	rl.CloseWindow()
}
