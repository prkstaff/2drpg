package scenes

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	OnLoad()
	Update()
	Draw(screen *ebiten.Image)
}
