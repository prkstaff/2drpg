package main

import (
	"bytes"
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
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
}

func (startScreen *GameStartScreen) onLoad() {
	// Parse .tmx file.
	gameMap, err := tiled.LoadFromFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
	startScreen.gameMap = gameMap
	// create a renderer
	mapRenderer, err := render.NewRenderer(gameMap)

	if err != nil {
		panic(err)
	}
	// render it to an in memory image
	err = mapRenderer.RenderVisibleLayers()

	if err != nil {
		panic(err)
	}

	var buff []byte
	buffer := bytes.NewBuffer(buff)

	mapRenderer.SaveAsPng(buffer)

	im, err := png.Decode(buffer)

	startScreen.mapBGImage = ebiten.NewImageFromImage(im)
}
