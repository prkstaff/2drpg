package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Scene interface {
	update()
	draw(screen *ebiten.Image)
}

type GameStartScreen struct {
	clock string
}

func (startScreen *GameStartScreen) update() {
	dt := time.Now()
	startScreen.clock = fmt.Sprintf("time now: %s", dt.String())
}

func (startScreen *GameStartScreen) draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, startScreen.clock)
}
