package scenes

import (
	"embed"
	"fmt"
	"image"
	"os"
	"time"

	"github.com/Rulox/ebitmx"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/prkstaff/2drpg/settings"
)

type VillageScene struct {
	clock           string
	gameMap         *ebitmx.EbitenMap
	mapBGImage      *ebiten.Image
	tilesetMetadata *ebitmx.EbitenTileset
	EmbeddedFS      *embed.FS
}

func (startScreen *VillageScene) getTileImgByID(id int) *ebiten.Image {
	// The tsx format starts counting tiles from 1, so to make these calculations
	// work correctly, we need to decrement the ID by 1
	id -= 1

	x0 := (id % startScreen.tilesetMetadata.TilesetWidth) * startScreen.tilesetMetadata.TileWidth
	y0 := (id / startScreen.tilesetMetadata.TilesetWidth) * startScreen.tilesetMetadata.TileHeight
	x1, y1 := x0+startScreen.tilesetMetadata.TileWidth, y0+startScreen.tilesetMetadata.TileHeight

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
	sx := float64(settings.GameSettings().ScreenWidth / (startScreen.gameMap.MapWidth * startScreen.gameMap.TileWidth))
	sy := float64(settings.GameSettings().ScreenHeight / (startScreen.gameMap.MapHeight * startScreen.gameMap.TileHeight))
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
	startScreen.gameMap, loadMapTMXErr = ebitmx.GetEbitenMapFromFS(startScreen.EmbeddedFS, "assets/maps/main.tmx")
	if loadMapTMXErr != nil {
		fmt.Println(loadMapTMXErr)
		os.Exit(2)
	}

	var tilesetLoadErr error
	startScreen.tilesetMetadata, tilesetLoadErr = ebitmx.GetTilesetFromFS(startScreen.EmbeddedFS, "assets/tilesets/tileset.tsx")
	if tilesetLoadErr != nil {
		fmt.Println(tilesetLoadErr)
		os.Exit(2)
	}
}
