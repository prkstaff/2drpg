package main

import (
	"fmt"
	_ "image/png"
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
	clock      string
	gameMap    *tiled.Map
	mapBGImage *ebiten.Image
}

func (startScreen *GameStartScreen) update() {
	dt := time.Now()
	startScreen.clock = fmt.Sprintf("time now: %s", dt.String())
}

func (startScreen *GameStartScreen) draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, startScreen.clock)
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(startScreen.mapBGImage, op)
}

func (startScreen *GameStartScreen) onLoad() {
	fmt.Println("Loaded")
	// Parse .tmx file.
	gameMap, err := tiled.LoadFromFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
	startScreen.gameMap = gameMap

	myBG, _, err2 := ebitenutil.NewImageFromFile("tileset.png")
	if err2 != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	startScreen.mapBGImage = myBG
}
