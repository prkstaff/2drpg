package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
)

const mapPath = "first.tmx"

type Scene interface {
	onLoad()
	update()
	draw(screen *ebiten.Image)
}

type GameStartScreen struct {
	clock   string
	gameMap *tiled.Map
}

func (startScreen *GameStartScreen) update() {
	dt := time.Now()
	startScreen.clock = fmt.Sprintf("time now: %s", dt.String())
}

func (startScreen *GameStartScreen) draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, startScreen.clock)
}

func (startScreen *GameStartScreen) onLoad() {
	// Parse .tmx file.
	gameMap, err := tiled.LoadFromFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
	startScreen.gameMap = gameMap
}
