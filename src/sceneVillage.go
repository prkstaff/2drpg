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
	inputHandler    InputHandler
	tilesetLayers   [][]int32
	tilesetObjectGroups map[int32][]*ObjectGroup
}

func (v VillageScene) getTileIDByIndex(index int32) []int32 {
	// will return any tileset collision box in any layer
	var tileIDs []int32
	for _, l := range v.tilesetLayers {
		tileIDs = append(tileIDs, l[index])
	}
	return tileIDs
}

func (v VillageScene) getCollisionObjectGroupByTileIndexes(indexes []int32) []Object {
	var objects []Object
	for _, index := range indexes {
		// get every tileid from every layer
		tileIDs := v.getTileIDByIndex(index)
		for _, tileID := range tileIDs {
			objGls, exists := v.tilesetObjectGroups[tileID]
			if exists {
				for _, objG := range objGls{
					for _, obj := range objG.Objects{
						valueObj := *obj
						valueObj.MapX = float64(index % v.gameMap.Width * int32(v.gameMap.Tileset.TileWidth))
						valueObj.MapY = float64(index / v.gameMap.Width * int32(v.gameMap.Tileset.TileHeight))
						objects = append(objects, valueObj)
					}
				}
			}
		}
	}
	return objects
}

func (v *VillageScene) HeroDontColideAgainsTileset(hero *Hero, orientation string) bool {
	heroY := float64(hero.YPos + hero.SpriteHeight / 2) // we want the collision only in the feet
	heroX := float64(hero.XPos)
	// first we get the x,y coordinates, based on x,y we get the current tile index
	heroTileIndex := v.getCurrentHeroTileIndex()
	//fmt.Printf("Current hero index is %v\n", heroTileIndex)
	if orientation == "up" {
		directUpperTileIndex := heroTileIndex - v.gameMap.Width
		directUpperTileRightAdjacentIndex := directUpperTileIndex+1
		directUpperTileLeftAdjacentIndex := directUpperTileIndex-1
		upperTiles := []int32{directUpperTileIndex, directUpperTileLeftAdjacentIndex, directUpperTileRightAdjacentIndex,
			heroTileIndex, heroTileIndex-1, heroTileIndex+1, heroTileIndex+v.gameMap.Width,
			heroTileIndex+v.gameMap.Width-1, heroTileIndex+v.gameMap.Width+1}
		colisionObjects := v.getCollisionObjectGroupByTileIndexes(upperTiles)
		if len(colisionObjects) > 0 {
			for _, obj := range colisionObjects {
				obX := obj.X + obj.MapX // gets the real X coordinate in map
				obY := obj.Y + obj.MapY // gets the real X coordinate in map
				obXEnd :=  obX + obj.Width
				obYEnd :=  obY + obj.Height
				//rect1.x < rect2.x + rect2.width && rect1.x + rect1.width > rect2.x
				fmt.Printf("Hero X: %v, Y: %x\n", heroX, heroY)
				if obX < (heroX + float64(hero.SpriteWidth)) && obXEnd > heroX {
					// check if hero feet is inside the box
					fmt.Println("Object X: %v, Y: %v\n", obX, obY)
					if heroY < obYEnd{
						return false
					}
				}
			}
		}
	}
	// if the hero is is heading up we will check for the adjacent tiles in the upper row.
	// if any of these tiles have collision box we will check for the collision
	return true
}
func (v VillageScene) getCurrentHeroTileIndex() int32 {
	hero := v.characters[0]
	currentColum := math.Floor(float64(hero.XPos) / float64(v.gameMap.Tileset.TileWidth))
	currentRow := math.Floor(float64(hero.YPos) / float64(v.gameMap.Tileset.TileHeight))
	return int32(currentRow*float64(v.gameMap.Width) + currentColum)
}
func (v *VillageScene) Update(keyStates []uint8) {
	hero := v.characters[0]

	v.inputHandler.Commands = nil
	var anyKeyMovePressed bool
	// quando uma tecla ?? solta e a outra mant??m pressionada, se perde a que estava pressionada.
	if keyStates[sdl.SCANCODE_W] == 1 && v.HeroDontColideAgainsTileset(hero, "up") {
		anyKeyMovePressed = true
		mvUp := MoveUpCommand{}
		v.inputHandler.Commands = append(v.inputHandler.Commands, mvUp)
	}
	if keyStates[sdl.SCANCODE_S] == 1 && v.HeroDontColideAgainsTileset(hero, "down") {
		anyKeyMovePressed = true
		mvDw := MoveDownCommand{}
		v.inputHandler.Commands = append(v.inputHandler.Commands, mvDw)
	}
	if keyStates[sdl.SCANCODE_A] == 1 && v.HeroDontColideAgainsTileset(hero, "left") {
		anyKeyMovePressed = true
		mvLf := MoveLeftCommand{}
		v.inputHandler.Commands = append(v.inputHandler.Commands, mvLf)
	}
	if keyStates[sdl.SCANCODE_D] == 1 && v.HeroDontColideAgainsTileset(hero, "right") {
		anyKeyMovePressed = true
		mvRg := MoveRightCommand{}
		v.inputHandler.Commands = append(v.inputHandler.Commands, mvRg)
	}
	v.inputHandler.HandleInput(hero)
	if anyKeyMovePressed {
		hero.MoveKeyPressed = true
	} else {
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

	for lId, l := range v.tilesetLayers {
		// Draw tiles
		for i, id := range l {
			id -= 1

			// Coordinates of sprite destination
			ix0 := (i % int(v.gameMap.Width)) * tileset.TileWidth
			iy0 := (i / int(v.gameMap.Width)) * tileset.TileHeight

			// scaled destinations
			scaledXPos := float64(ix0) * safeScale
			scaledYPos := float64(iy0) * safeScale

			src := tileset.GetTileRectSliceFromTilesetByID(id)
			dst := sdl.Rect{int32(scaledXPos), int32(scaledYPos), int32(16 * safeScale), int32(16 * safeScale)}
			err := renderer.Copy(v.gameMap.Tileset.Texture, &src, &dst)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		}
		// Draw Characters
		for _, obj := range v.characters {
			if int(obj.DrawAfterLayer) == lId {
				obj.Draw(renderer, int32(safeScale))
			}
		}
	}
}

func (v *VillageScene) OnLoad(renderer *sdl.Renderer) {
	v.inputHandler = InputHandler{}
	var loadMapTMXErr error
	v.gameMap, loadMapTMXErr = ReadTMX("assets/maps/main.tmx")
	_, err2 := v.gameMap.Tileset.LoadDataFromTSXFile()
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(2)
	}

	//Decode Tileset Layers
	for _, l := range v.gameMap.Layers {
		layerTilesIDSlice, err := l.Data.DecodeCSVTileData()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		v.tilesetLayers = append(v.tilesetLayers, layerTilesIDSlice)
	}

	// map tilseset collision objects by ID
	v.tilesetObjectGroups = make(map[int32][]*ObjectGroup)
	for _, tile := range v.gameMap.Tileset.Tiles{
		if len(tile.ObjectGroups) > 0 {
			v.tilesetObjectGroups[tile.ID] = tile.ObjectGroups
		}
	}

	startPosObj, startPosObjErr := v.gameMap.Objects.GetObjectByName("StartPos")
	if startPosObjErr != nil {
		log.Print(startPosObjErr)
		os.Exit(2)
	}
	hero := Hero{
		SpritePath:        "assets/tilesets/character.png",
		Texture:           nil,
		SpriteWidth:       16,
		SpriteHeight:      16,
		XPos:              int32(startPosObj.X),
		YPos:              int32(startPosObj.Y),
		DrawAfterLayer:    3,
		AnimationFrame:    0,
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
