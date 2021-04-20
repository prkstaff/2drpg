package characters

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Actor interface {
	MoveUp()
	MoveDown()
	MoveLeft()
	MoveRight()
}

type Hero struct {
	SpritePath     string
	Texture        *sdl.Texture
	SpriteWidth    uint16
	SpriteHeight   uint16
	XPos           uint16
	YPos           uint16
	DrawAfterLayer uint8
}

func (h *Hero) LoadSpriteIMG(renderer *sdl.Renderer) {
	surface, err := img.Load(h.SpritePath)
	if err != nil {
		fmt.Printf("Failed to load PNG: %s\n", err)
		os.Exit(2)
	}
	defer surface.Free()

	h.Texture, err = renderer.CreateTextureFromSurface(surface)
	if err != nil {
		fmt.Printf("Failed to create texture: %s\n", err)
		os.Exit(2)
	}
}

func (h Hero) Draw(renderer *sdl.Renderer, scale int32) {
	src := sdl.Rect{X: 0, Y: 0, W: 16, H: 32}
	dst := sdl.Rect{X: int32(h.XPos) * scale, Y: int32(h.YPos) * scale, W: 16 * scale, H: 32 * scale}
	renderer.Copy(h.Texture, &src, &dst)
}

func (h Hero) MoveUp() {
	fmt.Println("up")
}
func (h Hero) MoveDown() {
	fmt.Println("down")
}
func (h Hero) MoveLeft() {
	fmt.Println("left")
}
func (h Hero) MoveRight() {
	fmt.Println("right")
}
