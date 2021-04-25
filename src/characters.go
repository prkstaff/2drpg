package src

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
	SpritePath                 string
	SpriteOrientation          string
	AnimationFrame             int32 // 1 - 4
	Texture                    *sdl.Texture
	SpriteWidth                int32
	SpriteHeight               int32
	XPos                       int32
	YPos                       int32
	DrawAfterLayer             int32
	tileset                    Tileset
	MoveKeyPressed             bool
	TilesetTiles               TilesetTilesAWSD
	lastAnimationFrameRotation int32
	SpriteManager              SpriteManager
}

func (h *Hero) Load(renderer *sdl.Renderer) {
	h.LoadSpriteIMG(renderer)
	h.LoadTileset()
}
func (h *Hero) LoadTileset() {
	h.tileset = Tileset{Source: "heroSpriteSheet.tsx"}
	h.tileset.LoadDataFromTSXFile()
	var err error
	h.TilesetTiles.Up, err = GetTilesetTileByID(34, h.tileset)
	if err != nil {
		fmt.Println("failed getting tileset tile for up")
	}
	h.TilesetTiles.Down, err = GetTilesetTileByID(0, h.tileset)
	if err != nil {
		fmt.Println("failed getting tileset tile for down")
	}
	h.TilesetTiles.Left, err = GetTilesetTileByID(51, h.tileset)
	if err != nil {
		fmt.Println("failed getting tileset tile for left")
	}
	h.TilesetTiles.Right, err = GetTilesetTileByID(17, h.tileset)
	if err != nil {
		fmt.Println("failed getting tileset tile for right")
	}
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
	// return the rectangle for slice the spritesheet
	// Either for rendering animation of hero staying still
	var spriteRectCut sdl.Rect
	if h.SpriteOrientation == "up" {
		spriteRectCut = h.SpriteManager.GetRectSprite(h.MoveKeyPressed, h.TilesetTiles.Up, h.tileset)
	} else if h.SpriteOrientation == "down" {
		spriteRectCut = h.SpriteManager.GetRectSprite(h.MoveKeyPressed, h.TilesetTiles.Down, h.tileset)
	} else if h.SpriteOrientation == "left" {
		spriteRectCut = h.SpriteManager.GetRectSprite(h.MoveKeyPressed, h.TilesetTiles.Left, h.tileset)
	} else if h.SpriteOrientation == "right" {
		spriteRectCut = h.SpriteManager.GetRectSprite(h.MoveKeyPressed, h.TilesetTiles.Right, h.tileset)
	} else {
		fmt.Println("Unexpected orientation")
		os.Exit(2)
	}
	return spriteRectCut
}
func (h *Hero) Draw(renderer *sdl.Renderer, scale int32) {
	spriteRectCut := h.getHeroRectangleForSpriteRender()
	dst := sdl.Rect{X: int32(h.XPos) * scale, Y: int32(h.YPos) * scale, W: 16 * scale, H: 32 * scale}
	renderer.Copy(h.Texture, &spriteRectCut, &dst)
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
