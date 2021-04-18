package characters

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type Hero struct{
	SpritePath string
	Texture *sdl.Texture
	SpriteWidth uint16
	SpriteHeight uint16
	XPos uint16
	YPos uint16
	DrawAfterLayer uint8
}

func (h *Hero) LoadSpriteIMG(renderer *sdl.Renderer)  {
	surface, err := img.Load(h.SpritePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		os.Exit(2)
	}
	defer surface.Free()

	h.Texture, err = renderer.CreateTextureFromSurface(surface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(2)
	}
}

func (h Hero) Draw(renderer *sdl.Renderer, scale int32) {
	src := sdl.Rect{X: 0, Y:0, W:16, H:32}
	dst := sdl.Rect{X: int32(h.XPos)*scale, Y:int32(h.YPos)*scale, W:16*scale, H:32*scale}
	renderer.Copy(h.Texture, &src, &dst)
}