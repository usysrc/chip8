package emulator

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	DotSize = int32(8)
	Width   = 64
	Height  = 32
)

type Display struct {
	pixels [Width * Height]bool
}

func (d *Display) Init() error {
	rl.InitWindow(800, 450, "chip8")
	rl.SetTargetFPS(60)
	return nil
}

func (d *Display) Clear() {
	for i := 0; i < Width; i++ {
		for j := 0; j < Height; j++ {
			d.pixels[j*Width+i] = false
		}
	}
}

func (d *Display) Set(x, y uint16) error {
	if x < 0 || x > Width || y < 0 || y > Height {
		return fmt.Errorf("trying to draw out of screen bounds")
	}
	d.pixels[y*Width+x] = true
	return nil
}

func (d *Display) Draw() error {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	for i := 0; i < Width; i++ {
		for j := 0; j < Height; j++ {
			if d.pixels[j*Width+i] {
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
