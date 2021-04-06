package main

import (
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"os"
	"time"

	"github.com/Rulox/ebitmx"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// embeddedFS holds our game assets so we can distribute our game as a single binary
//go:embed first.tmx default-tileset.tsx tileset.png
var embeddedFS embed.FS

const (
	screenWidth  = 960
	screenHeight = 960
)

type Scene interface {
	onLoad()
	update()
	draw(screen *ebiten.Image)
}

type GameStartScreen struct {
	clock           string
	gameMap         *ebitmx.EbitenMap
	mapBGImage      *ebiten.Image
	tilesetMetadata *ebitmx.EbitenTileset
}

func (startScreen *GameStartScreen) getTileImgByID(id int) *ebiten.Image {
	// The tsx format starts counting tiles from 1, so to make these calculations
	// work correctly, we need to decrement the ID by 1
	id -= 1

	x0 := (id % startScreen.tilesetMetadata.TilesetWidth) * startScreen.tilesetMetadata.TileWidth
	y0 := (id / startScreen.tilesetMetadata.TilesetWidth) * startScreen.tilesetMetadata.TileHeight
	x1, y1 := x0+startScreen.tilesetMetadata.TileWidth, y0+startScreen.tilesetMetadata.TileHeight

	return startScreen.mapBGImage.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
}

func (startScreen *GameStartScreen) update() {
	dt := time.Now()
	startScreen.clock = fmt.Sprintf("time now: %s", dt.String())
}

func (startScreen *GameStartScreen) draw(screen *ebiten.Image) {
	// Draw map using the same method as the official tiles example
	// https://ebiten.org/examples/tiles.html

	// The scaling we use is consistent across all tiles, so we'll
	// calculate it outside of the tile-drawing loop
	sx := float64(screenWidth / (startScreen.gameMap.MapWidth * startScreen.gameMap.TileWidth))
	sy := float64(screenHeight / (startScreen.gameMap.MapHeight * startScreen.gameMap.TileHeight))
	for _, l := range startScreen.gameMap.Layers {
		for i, id := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(
				float64((i%startScreen.gameMap.MapWidth)*startScreen.gameMap.TileWidth),
				float64((i/startScreen.gameMap.MapHeight)*startScreen.gameMap.TileHeight),
			)
			op.GeoM.Scale(sx, sy)

			screen.DrawImage(startScreen.getTileImgByID(id), op)
		}
	}

	ebitenutil.DebugPrint(screen, startScreen.clock)
}

func (startScreen *GameStartScreen) onLoad() {
	var tilesetIMG *ebiten.Image
	{
		imgFile, err := embeddedFS.Open("tileset.png")
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		img, _, err := image.Decode(imgFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		tilesetIMG = ebiten.NewImageFromImage(img)
	}
	startScreen.mapBGImage = tilesetIMG
	var loadMapTMXErr error
	startScreen.gameMap, loadMapTMXErr = ebitmx.GetEbitenMapFromFS(embeddedFS, "first.tmx")
	if loadMapTMXErr != nil {
		fmt.Println(loadMapTMXErr)
		os.Exit(2)
	}

	var tilesetLoadErr error
	startScreen.tilesetMetadata, tilesetLoadErr = ebitmx.GetTilesetFromFS(embeddedFS, "default-tileset.tsx")
	if tilesetLoadErr != nil {
		fmt.Println(tilesetLoadErr)
		os.Exit(2)
	}
}
