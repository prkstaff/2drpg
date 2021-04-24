package src

import (
	"embed"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type Scene interface {
	OnLoad(renderer *sdl.Renderer)
	Update(keyStates []uint8)
	Draw(renderer *sdl.Renderer)
}

type VillageScene struct {
	gameMap         *Map
	EmbeddedFS      *embed.FS
	characters      []*Hero
	initialHeroPosX uint16
	initialHeroPosY uint16
	inputHandler    InputHandler
}

func (v *VillageScene) Update(keyStates []uint8) {
	hero := v.characters[0]

	v.inputHandler.Commands = nil
	var anyKeyMovePressed bool
	// quando uma tecla é solta e a outra mantém pressionada, se perde a que estava pressionada.
	if keyStates[sdl.SCANCODE_W] == 1 && v.gameMap.Tileset.HeroDontColideAgainsTileset("up"){
		anyKeyMovePressed = true
		mvUp := MoveUpCommand{}
		v.inputHandler.Commands = append(v.inputHandler.Commands, mvUp)
	}
	if keyStates[sdl.SCANCODE_S] == 1 && v.gameMap.Tileset.HeroDontColideAgainsTileset("down"){
		anyKeyMovePressed = true
		mvDw := MoveDownCommand{}
		v.inputHandler.Commands = append(v.inputHandler.Commands, mvDw)
	}
	if keyStates[sdl.SCANCODE_A] == 1 && v.gameMap.Tileset.HeroDontColideAgainsTileset("left"){
		anyKeyMovePressed = true
		mvLf := MoveLeftCommand{}
		v.inputHandler.Commands = append(v.inputHandler.Commands, mvLf)
	}
	if keyStates[sdl.SCANCODE_D] == 1 && v.gameMap.Tileset.HeroDontColideAgainsTileset("right"){
		anyKeyMovePressed = true
		mvRg := MoveRightCommand{}
		v.inputHandler.Commands = append(v.inputHandler.Commands, mvRg)
	}
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
	wWidth := GameSettings().WindowWidth
	wHeight := GameSettings().WindowHeigh
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

			// Coordinates of sprite destination
			ix0 := (i % int(v.gameMap.Width)) * tileset.TileWidth
			iy0 := (i / int(v.gameMap.Width)) * tileset.TileHeight

			// scaled destinations
			scaledXPos := float64(ix0) * safeScale
			scaledYPos := float64(iy0) * safeScale

			src := tileset.GetTileRectSliceFromTilesetByID(id)
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
	v.inputHandler = InputHandler{}
	var loadMapTMXErr error
	v.gameMap, loadMapTMXErr = ReadTMX("assets/maps/main.tmx")
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
	hero := Hero{
		SpritePath:     "assets/tilesets/character.png",
		Texture:        nil,
		SpriteWidth:    16,
		SpriteHeight:   16,
		XPos:           uint16(startPosObj.X),
		YPos:           uint16(startPosObj.Y),
		DrawAfterLayer: 3,
		AnimationFrame: 0,
		SpriteOrientation: "down",
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
