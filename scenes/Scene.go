package scenes

import (
	_ "image/png"

	"github.com/veandco/go-sdl2/sdl"
)

type Scene interface {
	OnLoad(renderer *sdl.Renderer)
	Update(keyStates []uint8)
	Draw(renderer *sdl.Renderer)
}
