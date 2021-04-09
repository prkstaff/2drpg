package tiled

import "testing"

func TestReadTMX(t *testing.T) {
	tiledMap := ReadTMX("mainTest.tmx")
	if tiledMap.Width != 50 || tiledMap.Height != 50 {
		t.Error("unexpected value for map width or length")
	}
}
