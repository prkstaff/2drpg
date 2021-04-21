package sprite

import (
	"errors"
	"fmt"
	"github.com/prkstaff/2drpg/tiled"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

// TileAnimationKeys the animation tile ids from spritesheet tileset
type TilesetTilesAWSD struct {
	Up *tiled.TilesetTile
	Down *tiled.TilesetTile
	Left *tiled.TilesetTile
	Right *tiled.TilesetTile
}

type SpriteManager struct {
	CurrentTileId int32
	TimeAppliedLastFrame time.Time
	LastTileId int32
	LastFrameTimeDuration int32 // the time the tile should stay until next tile frame
	shouldAnimate bool
	currentFrame int32
}

func GetTilesetTileByID(tileID int32, tileset tiled.Tileset) (*tiled.TilesetTile, error) {
	for _, tile := range tileset.Tiles{
		if tile.ID == tileID{
			return tile, nil
		}
	}
	return nil, errors.New("not found tileid in tile")
}

func (s *SpriteManager) GetRectSprite(shouldAnimate bool, tile *tiled.TilesetTile, tileset tiled.Tileset) sdl.Rect{
	// tiled produces a tileset tsx file
	// the tileID is the Tile Id that have an animation and is the first Tile Frame ID
	// should animate will say if it should rotate between the frames or stay with the first tile still.
	var rect sdl.Rect
	frames := tile.Animation
	if shouldAnimate {
		rect = sdl.Rect{X: 96, Y: 32, W: 16, H: 32}
	}else{
		s.LastTileId = tile.ID
		s.TimeAppliedLastFrame = time.Now()
		s.currentFrame = 0

		if tile.ID == 34{
			fmt.Println("b")
		}
		rect = tileset.GetTileRectSliceFromTilesetByID(frames[0].TileID)
	}
	return rect
}