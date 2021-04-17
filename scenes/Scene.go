package scenes

import (
	"github.com/veandco/go-sdl2/sdl"
	_ "image/png"
)

type Scene interface {
	OnLoad(renderer *sdl.Renderer)
	Update()
	Draw(renderer *sdl.Renderer)
}
