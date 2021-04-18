package scenes

import (
	"embed"
	"fmt"
	"github.com/prkstaff/2drpg/characters"
	"github.com/prkstaff/2drpg/tiled"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
	"time"
)

type VillageScene struct {
	clock           string
	gameMap         *tiled.Map
	EmbeddedFS      *embed.FS
	characters []characters.Hero
	initialHeroPosX uint16
	initialHeroPosY uint16
}

func (startScreen *VillageScene) Update() {
	dt := time.Now()
	startScreen.clock = fmt.Sprintf("time now: %s", dt.String())
}

func (startScreen *VillageScene) Draw(renderer *sdl.Renderer) {
	//sx := float64(settings.GameSettings().ScreenWidth / (startScreen.gameMap.Width * startScreen.gameMap.TileWidth))
	//sy := float64(settings.GameSettings().ScreenHeight / (startScreen.gameMap.Height * startScreen.gameMap.TileHeight))
	for _, l := range startScreen.gameMap.Layers {
		layerTilesIDSlice, err := l.Data.DecodeCSVTileData()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		// Draw tiles
		for i, id := range layerTilesIDSlice {
			id -= 1

			// Coordinates of sprite slice
			x0 := (int(id) % startScreen.gameMap.Tileset.Columns) * startScreen.gameMap.Tileset.TileWidth
			y0 := (int(id) / startScreen.gameMap.Tileset.Columns) * startScreen.gameMap.Tileset.TileHeight
			//x1, y1 := x0+startScreen.gameMap.Tileset.TileWidth, y0+startScreen.gameMap.Tileset.TileHeight

			// Coordinates of sprite destination
			ix0 := (i % int(startScreen.gameMap.Width)) * startScreen.gameMap.Tileset.TileWidth
			iy0 := (i / int(startScreen.gameMap.Width)) * startScreen.gameMap.Tileset.TileHeight

			// Calculate draw scale
			//tileseRawTotalWidth := startScreen.gameMap.Tileset.Columns * startScreen.gameMap.Tileset.TileWidth

			src := sdl.Rect{int32(x0), int32(y0), 16, 16}
			dst := sdl.Rect{int32(ix0), int32(iy0), 16, 16}
			renderer.Copy(startScreen.gameMap.Tileset.Texture, &src, &dst)
		}
	}
	//	// Draw Characters
	//	for _, obj := range startScreen.characters {
	//		if uint32(obj.DrawAfterLayer) == l.ID {
	//			obj.Draw(screen, *startScreen.gameMap)
	//		}
	//	}
	//}
	//https://github.com/lafriks/go-tiled

	//ebitenutil.DebugPrint(screen, startScreen.clock)
}

func (startScreen *VillageScene) OnLoad(renderer *sdl.Renderer) {
	var loadMapTMXErr error
	startScreen.gameMap, loadMapTMXErr = tiled.ReadTMX("assets/maps/main.tmx")
	_, err2 := startScreen.gameMap.Tileset.LoadDataFromTSXFile()
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(2)
	}

	startPosObj, startPosObjErr := startScreen.gameMap.Objects.GetObjectByName("StartPos")
	if startPosObjErr != nil {
		log.Print(startPosObjErr)
		os.Exit(2)
	}
	hero := characters.Hero{
		SpritePath:   "assets/tilesets/character.png",
		Sprite:       nil,
		SpriteWidth:  16,
		SpriteHeight: 16,
		XPos:         uint16(startPosObj.X),
		YPos:         uint16(startPosObj.Y),
		DrawAfterLayer: 2,
	}
	hero.LoadSpriteIMG()
	startScreen.characters = append(startScreen.characters, hero)
	startScreen.gameMap.Tileset.LoadTileSetTexture(renderer)
	if loadMapTMXErr != nil {
		fmt.Println(loadMapTMXErr)
		os.Exit(2)
	}
}
func (startScreen *VillageScene) onDestroy() {
	defer startScreen.gameMap.Tileset.Texture.Destroy()
}
