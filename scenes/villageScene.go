package scenes

import (
	"embed"
	"fmt"
	"github.com/prkstaff/2drpg/characters"
	"github.com/prkstaff/2drpg/settings"
	"github.com/prkstaff/2drpg/tiled"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"math"
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
	tileset := startScreen.gameMap.Tileset
	// Calculate draw scale
	wWidth := settings.GameSettings().WindowWidth
	wHeight := settings.GameSettings().WindowHeigh
	sx := float64(wWidth) / (float64(startScreen.gameMap.Width) * float64(startScreen.gameMap.TileWidth))
	sy := float64(wHeight) / (float64(startScreen.gameMap.Height) * float64(startScreen.gameMap.TileHeight))

	var safeScale float64

	// since its not good to strech a pixel by subdecimal points we will floor
	// also to fit the screen, ceil would not fit the screen
	if sx < sy {
		safeScale = math.Floor(sx)
	} else {
		safeScale = math.Floor(sy)
	}
	safeScale += 1

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
			x0 := (int(id) % tileset.Columns) * tileset.TileWidth
			y0 := (int(id) / tileset.Columns) * tileset.TileHeight
			//x1, y1 := x0+startScreen.gameMap.Tileset.TileWidth, y0+startScreen.gameMap.Tileset.TileHeight

			// Coordinates of sprite destination
			ix0 := (i % int(startScreen.gameMap.Width)) * tileset.TileWidth
			iy0 := (i / int(startScreen.gameMap.Width)) * tileset.TileHeight



			// scaled destinations
			scaledXPos := float64(ix0) * safeScale
			scaledYPos := float64(iy0) * safeScale

			src := sdl.Rect{int32(x0), int32(y0), 16, 16}
			dst := sdl.Rect{int32(scaledXPos), int32(scaledYPos), int32(16 * safeScale), int32(16 * safeScale)}
			renderer.Copy(startScreen.gameMap.Tileset.Texture, &src, &dst)
		}
		// Draw Characters
		for _, obj := range startScreen.characters {
			if uint32(obj.DrawAfterLayer) == l.ID {
				obj.Draw(renderer, int32(safeScale))
			}
		}
	}
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
		Texture:       nil,
		SpriteWidth:  16,
		SpriteHeight: 16,
		XPos:         uint16(startPosObj.X),
		YPos:         uint16(startPosObj.Y),
		DrawAfterLayer: 2,
	}
	hero.LoadSpriteIMG(renderer)
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
