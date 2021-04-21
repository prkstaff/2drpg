package scenes

import (
	"embed"
	"fmt"
	"github.com/prkstaff/2drpg/sprite"
	"log"
	"math"
	"os"

	"github.com/prkstaff/2drpg/characters"
	"github.com/prkstaff/2drpg/input"
	"github.com/prkstaff/2drpg/settings"
	"github.com/prkstaff/2drpg/tiled"
	"github.com/veandco/go-sdl2/sdl"
)

type VillageScene struct {
	gameMap         *tiled.Map
	EmbeddedFS      *embed.FS
	characters      []*characters.Hero
	initialHeroPosX uint16
	initialHeroPosY uint16
	inputHandler    input.InputHandler
}

func (v *VillageScene) Update(keyStates []uint8) {
	v.inputHandler.Commands = nil
	var anyKeyMovePressed bool
	if keyStates[sdl.SCANCODE_W] == 1 {
		anyKeyMovePressed = true
		mvUp := input.MoveUpCommand{}
		v.inputHandler.Commands = append(v.inputHandler.Commands, mvUp)
	}
	if keyStates[sdl.SCANCODE_S] == 1 {
		anyKeyMovePressed = true
		mvDw := input.MoveDownCommand{}
		v.inputHandler.Commands = append(v.inputHandler.Commands, mvDw)
	}
	if keyStates[sdl.SCANCODE_A] == 1 {
		anyKeyMovePressed = true
		mvLf := input.MoveLeftCommand{}
		v.inputHandler.Commands = append(v.inputHandler.Commands, mvLf)
	}
	if keyStates[sdl.SCANCODE_D] == 1 {
		anyKeyMovePressed = true
		mvRg := input.MoveRightCommand{}
		v.inputHandler.Commands = append(v.inputHandler.Commands, mvRg)
	}
	hero := v.characters[0]
	v.inputHandler.HandleInput(hero)
	if anyKeyMovePressed {
		hero.MoveKeyPressed = true
	}else{
		hero.MoveKeyPressed = false
	}
}

func (v *VillageScene) Draw(renderer *sdl.Renderer) {
	tileset := v.gameMap.Tileset
	// Calculate draw scale
	wWidth := settings.GameSettings().WindowWidth
	wHeight := settings.GameSettings().WindowHeigh
	sx := float64(wWidth) / (float64(v.gameMap.Width) * float64(v.gameMap.TileWidth))
	sy := float64(wHeight) / (float64(v.gameMap.Height) * float64(v.gameMap.TileHeight))

	var safeScale float64

	// since its not good to strech a pixel by subdecimal points we will floor
	// also to fit the screen, ceil would not fit the screen
	if sx < sy {
		safeScale = math.Floor(sx)
	} else {
		safeScale = math.Floor(sy)
	}
	safeScale += 1

	for _, l := range v.gameMap.Layers {
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
			ix0 := (i % int(v.gameMap.Width)) * tileset.TileWidth
			iy0 := (i / int(v.gameMap.Width)) * tileset.TileHeight

			// scaled destinations
			scaledXPos := float64(ix0) * safeScale
			scaledYPos := float64(iy0) * safeScale

			src := sdl.Rect{int32(x0), int32(y0), 16, 16}
			dst := sdl.Rect{int32(scaledXPos), int32(scaledYPos), int32(16 * safeScale), int32(16 * safeScale)}
			renderer.Copy(v.gameMap.Tileset.Texture, &src, &dst)
		}
		// Draw Characters
		for _, obj := range v.characters {
			if uint32(obj.DrawAfterLayer) == l.ID {
				obj.Draw(renderer, int32(safeScale))
			}
		}
	}
}

func (v *VillageScene) OnLoad(renderer *sdl.Renderer) {
	fmt.Println("Called")
	v.inputHandler = input.InputHandler{}
	var loadMapTMXErr error
	v.gameMap, loadMapTMXErr = tiled.ReadTMX("assets/maps/main.tmx")
	_, err2 := v.gameMap.Tileset.LoadDataFromTSXFile()
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(2)
	}

	startPosObj, startPosObjErr := v.gameMap.Objects.GetObjectByName("StartPos")
	if startPosObjErr != nil {
		log.Print(startPosObjErr)
		os.Exit(2)
	}
	hero := characters.Hero{
		SpritePath:     "assets/tilesets/character.png",
		Texture:        nil,
		SpriteWidth:    16,
		SpriteHeight:   16,
		XPos:           uint16(startPosObj.X),
		YPos:           uint16(startPosObj.Y),
		DrawAfterLayer: 3,
		AnimationFrame: 0,
		SpriteOrientation: "down",
		TileKeys: sprite.TileAnimationKeys{Up:32,Down: 0,Left: 51,Right: 17},
	}
	hero.Load(renderer)
	v.characters = append(v.characters, &hero)
	v.gameMap.Tileset.LoadTileSetTexture(renderer)
	if loadMapTMXErr != nil {
		fmt.Println(loadMapTMXErr)
		os.Exit(2)
	}
}
func (v *VillageScene) onDestroy() {
	defer v.gameMap.Tileset.Texture.Destroy()
}
