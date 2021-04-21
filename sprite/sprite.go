package sprite

import (
	"github.com/prkstaff/2drpg/tiled"
	"github.com/veandco/go-sdl2/sdl"
)

// TileAnimationKeys the animation tile ids from spritesheet tileset
type TileAnimationKeys struct {
	Up int32
	Down int32
	Left int32
	Right int32
}

type SpriteManager struct {
	CurrentTileId int32
	shouldAnimate bool
	currentFrame int32
}

func (s *SpriteManager) GetRectSprite(shouldAnimate bool, tileId int32, tileset tiled.Tileset) sdl.Rect{
	// tiled produces a tileset tsx file
	// the tileID is the Tile Id that have an animation and is the first Tile Frame ID
	// should animate will say if it should rotate between the frames or stay with the first tile still.
	var rect sdl.Rect
	if shouldAnimate {
		rect = sdl.Rect{X: 96, Y: 32, W: 16, H: 32}
	}else{
		rect = tileset.GetTileRectSliceFromTilesetByID(tileId)
	}
	return rect
}