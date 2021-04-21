package characters

import (
	"fmt"
	"github.com/prkstaff/2drpg/sprite"
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
	SpritePath                 string
	SpriteOrientation          string
	AnimationFrame             int32 // 1 - 4
	Texture                    *sdl.Texture
	SpriteWidth                uint16
	SpriteHeight               uint16
	XPos                       uint16
	YPos                       uint16
	DrawAfterLayer             uint8
	tileset                    tiled.Tileset
	MoveKeyPressed             bool
	TileKeys   sprite.TileAnimationKeys
	lastAnimationFrameRotation int32
}

func (h *Hero) Load(renderer *sdl.Renderer) {
	h.LoadSpriteIMG(renderer)
	h.LoadTileset()
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
func (h *Hero) getHeroRectangleForSpriteRender() sdl.Rect {
	upRec := sdl.Rect{X: 0, Y: 64, W: 16, H: 32}
	downRec := sdl.Rect{X: 0, Y: 0, W: 16, H: 32}
	leftRec := sdl.Rect{X: 0, Y: 96, W: 16, H: 32}
	rightRec := sdl.Rect{X: 0, Y: 32, W: 16, H: 32}
	// return the rectangle for slice the spritesheet
	// Either for rendering animation of hero staying still
	var spriteRectCut sdl.Rect
	if h.MoveKeyPressed {
		spriteRectCut = sdl.Rect{X: 96, Y: 32, W: 16, H: 32}
		// animate with frame rotation
	} else {
		if h.SpriteOrientation == "up" {
			spriteRectCut = upRec
		} else if h.SpriteOrientation == "down" {
			spriteRectCut = downRec
		} else if h.SpriteOrientation == "left" {
			spriteRectCut = leftRec
		} else if h.SpriteOrientation == "right" {
			spriteRectCut = rightRec
		} else {
			fmt.Println("Unexpected orientation")
			os.Exit(2)
		}
	}
	return spriteRectCut
}
func (h *Hero) Draw(renderer *sdl.Renderer, scale int32) {

	spriteRectCut := h.getHeroRectangleForSpriteRender()
	dst := sdl.Rect{X: int32(h.XPos) * scale, Y: int32(h.YPos) * scale, W: 16 * scale, H: 32 * scale}
	renderer.Copy(h.Texture, &spriteRectCut, &dst)
	if h.AnimationFrame <= 3 {
		h.AnimationFrame += 16
	} else {
		h.AnimationFrame = 0
	}
	h.AnimationFrame += 1
}

func (h *Hero) MoveUp() {
	h.SpriteOrientation = "up"
	if h.YPos >= 1 {
		h.YPos = h.YPos - 1
	}
}
func (h *Hero) MoveDown() {
	h.SpriteOrientation = "down"
	h.YPos += 1
}
func (h *Hero) MoveLeft() {
	h.SpriteOrientation = "left"
	if h.XPos >= 1 {
		h.XPos -= 1
	}
}
func (h *Hero) MoveRight() {
	h.SpriteOrientation = "right"
	h.XPos += 1
}
