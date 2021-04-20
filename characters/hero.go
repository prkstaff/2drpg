package characters

import (
	"fmt"
	"os"

	"github.com/prkstaff/2drpg/tiled"
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
	tileset        tiled.Tileset
}

func (h *Hero) LoadTileset() {
	h.tileset = tiled.Tileset{Source: "heroSpriteSheet.tsx"}
	h.tileset.LoadDataFromTSXFile()
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

func (h *Hero) MoveUp() {
	if h.YPos >= 1 {
		h.YPos = h.YPos - 1
	}
}
func (h *Hero) MoveDown() {
	h.YPos += 1
}
func (h *Hero) MoveLeft() {
	if h.XPos >= 1 {
		h.XPos -= 1
	}
}
func (h *Hero) MoveRight() {
	h.XPos += 1
}
