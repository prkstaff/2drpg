package src

import (
	"errors"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

// TileAnimationKeys the animation tile ids from spritesheet tileset
type TilesetTilesAWSD struct {
	Up *TilesetTile
	Down *TilesetTile
	Left *TilesetTile
	Right *TilesetTile
}

type SpriteManager struct {
	CurrentTileId int32
	LastTimeFrameRotated time.Time
	LastTileId int32
	shouldAnimate bool
	currentFrame int32
}

func GetTilesetTileByID(tileID int32, tileset Tileset) (*TilesetTile, error) {
	for _, tile := range tileset.Tiles{
		if tile.ID == tileID{
			return tile, nil
		}
	}
	return nil, errors.New("not found tileid in tile")
}

func (s *SpriteManager) GetRectSprite(shouldAnimate bool, tile *TilesetTile, tileset Tileset) sdl.Rect{
	// tiled produces a tileset tsx file
	// the tileID is the Tile Id that have an animation and is the first Tile Frame ID
	// should animate will say if it should rotate between the frames or stay with the first tile still.
	var rect sdl.Rect
	frames := tile.Animation
	if shouldAnimate {
		elapsed := time.Since(s.LastTimeFrameRotated).Milliseconds()
		if s.LastTileId == tile.ID {
			// walking hero keep walking same direction, lets animate the framess
			if elapsed > int64(frames[s.currentFrame].Duration){
				if int(s.currentFrame) < len(frames)-1{
				// change the frame only if frame duration was respected
					s.currentFrame += 1
					s.LastTimeFrameRotated = time.Now()
				}else{
					s.currentFrame = 0
				}
			}
			rect = tileset.GetTileRectSliceFromTilesetByID(frames[s.currentFrame].TileID)
		}else{
			// walking hero is changing direction, lets just render the first frame, next will run in condition above
			s.LastTileId = tile.ID
			s.currentFrame = 0
			s.LastTimeFrameRotated = time.Now()
			rect = tileset.GetTileRectSliceFromTilesetByID(frames[s.currentFrame].TileID)
		}
	}else{
		s.LastTileId = tile.ID
		s.currentFrame = 0

		s.LastTimeFrameRotated = time.Now()
		rect = tileset.GetTileRectSliceFromTilesetByID(frames[0].TileID)
	}
	return rect
}