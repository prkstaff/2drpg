package scenes

import (
	"embed"
	"fmt"
	"github.com/prkstaff/2drpg/tiled"
	"image"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/prkstaff/2drpg/settings"
)

type VillageScene struct {
	clock           string
	gameMap         *tiled.Map
	mapBGImage      *ebiten.Image
	EmbeddedFS      *embed.FS
}

func (startScreen *VillageScene) getTileImgByID(id int) *ebiten.Image {
	// The tsx format starts counting tiles from 1, so to make these calculations
	// work correctly, we need to decrement the ID by 1
	id -= 1

	x0 := (id % startScreen.gameMap.Tileset.Columns) * startScreen.gameMap.Tileset.TileWidth
	y0 := (id / startScreen.gameMap.Tileset.Columns) * startScreen.gameMap.Tileset.TileHeight
	x1, y1 := x0+startScreen.gameMap.Tileset.TileWidth, y0+startScreen.gameMap.Tileset.TileHeight

	return startScreen.mapBGImage.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
}

func (startScreen *VillageScene) Update() {
	dt := time.Now()
	startScreen.clock = fmt.Sprintf("time now: %s", dt.String())
}

func (startScreen *VillageScene) Draw(screen *ebiten.Image) {
	// Draw map using the same method as the official tiles example
	// https://ebiten.org/examples/tiles.html

	// The scaling we use is consistent across all tiles, so we'll
	// calculate it outside of the tile-drawing loop
	sx := float64(settings.GameSettings().ScreenWidth / (startScreen.gameMap.Width * startScreen.gameMap.TileWidth))
	sy := float64(settings.GameSettings().ScreenHeight / (startScreen.gameMap.Height * startScreen.gameMap.TileHeight))
	for _, l := range startScreen.gameMap.Layers {
		layerTilesIDSlice, err := l.Data.DecodeCSVTileData()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		for i, id := range layerTilesIDSlice {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(
				float64((uint16(i)%startScreen.gameMap.Width)*startScreen.gameMap.TileWidth),
				float64((uint16(i)/startScreen.gameMap.Height)*startScreen.gameMap.TileHeight),
			)
			op.GeoM.Scale(sx, sy)

			screen.DrawImage(startScreen.getTileImgByID(int(id)), op)
		}
	}
	//https://github.com/lafriks/go-tiled

	//ebitenutil.DebugPrint(screen, startScreen.clock)
}

func (startScreen *VillageScene) OnLoad() {
	var tilesetIMG *ebiten.Image
	{
		imgFile, err := startScreen.EmbeddedFS.Open("assets/tilesets/tileset.png")
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
	startScreen.gameMap, loadMapTMXErr = tiled.ReadTMX("assets/maps/main.tmx")
	_, err2 := startScreen.gameMap.Tileset.LoadDataFromTSXFile()
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(2)
	}
	if loadMapTMXErr != nil {
		fmt.Println(loadMapTMXErr)
		os.Exit(2)
	}
}
