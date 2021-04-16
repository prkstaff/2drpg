package scenes

import (
	"github.com/veandco/go-sdl2/sdl"
	_ "image/png"
)

type Scene interface {
	OnLoad()
	Update()
	Draw(renderer *sdl.Renderer)
}
