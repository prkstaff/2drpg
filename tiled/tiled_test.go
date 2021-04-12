package tiled

import (
	"testing"
)

func TestReadTMX(t *testing.T) {
	tiledMap, _ := ReadTMX("mainTest.tmx")
	if tiledMap.Width != 50 || tiledMap.Height != 50 {
		t.Error("unexpected value for map width or length")
	}
	if tiledMap.TileHeight != 16 || tiledMap.TileWidth != 16 {
		t.Error("unexpected value for Tilewidth and/or tileHeight")
	}
	if tiledMap.Layers[0].Name != "Ground"{
		t.Error("Unexpected Name for first layer")
	}
	if tiledMap.Layers[0].Width != 50 || tiledMap.Layers[0].Height != 50{
		t.Error("Unexpected Widh or Heigh for first layer")
	}
	if tiledMap.Tileset.Source != "../tilesets/tileset.tsx"{
		t.Error("Unexpected source for tileset")
	}
	tiles, _ := tiledMap.Layers[0].Data.DecodeCSVTileData()
	// https://gamedevelopment.tutsplus.com/tutorials/parsing-and-rendering-tiled-tmx-format-maps-in-your-own-game-engine--gamedev-3104
}
