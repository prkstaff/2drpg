package tiled

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Image struct {
	XMLName xml.Name `xml:"image"`
	Format  string   `xml:"format"`
	Source  string   `xml:"source,attr"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
}

// AnimationFrame is single frame of animation
type AnimationFrame struct {
	// The local ID of a tile within the parent tileset.
	TileID uint32 `xml:"tileid,attr"`
	// How long (in milliseconds) this frame should be displayed before advancing to the next frame.
	Duration uint32 `xml:"duration,attr"`
}

// Properties wraps any number of custom properties
type Properties []*Property

// Property is used for custom properties
type Property struct {
	// The name of the property.
	Name string `xml:"name,attr"`
	// The type of the property. Can be string (default), int, float, bool, color or file (since 0.16, with color and file added in 0.17).
	Type string `xml:"type,attr"`
	// The value of the property.
	// Boolean properties have a value of either "true" or "false".
	// Color properties are stored in the format #AARRGGBB.
	// File properties are stored as paths relative from the location of the map file.
	Value string `xml:"value,attr"`
}

type TilesetTile struct {
	// The local tile ID within its tileset.
	ID int32 `xml:"id,attr"`
	// The type of the tile. Refers to an object type and is used by tile objects. (optional) (since 1.0)
	Type string `xml:"type,attr"`
	// Defines the terrain type of each corner of the tile, given as comma-separated indexes in the terrain types
	// array in the order top-left, top-right, bottom-left, bottom-right.
	// Leaving out a value means that corner has no terrain. (optional) (since 0.9)
	Terrain string `xml:"terrain,attr"`
	// A percentage indicating the probability that this tile is chosen when it competes with others while editing with the terrain tool. (optional) (since 0.9)
	Probability float32 `xml:"probability,attr"`
	// Custom properties
	Properties Properties `xml:"properties>property"`
	// Embedded image
	Image *Image `xml:"image"`
	// Tile object groups
	ObjectGroups []*ObjectGroup `xml:"objectgroup"`
	// List of animation frames
	Animation []*AnimationFrame `xml:"animation>frame"`
}

type Tileset struct {
	Source        string         `xml:"source,attr"`
	XMLName       xml.Name       `xml:"tileset"`
	Version       string         `xml:"version,attr"`
	Name          string         `xml:"name,attr"`
	TileWidth     int            `xml:"tilewidth,attr"`
	TileHeight    int            `xml:"tileheight,attr"`
	TileCount     int            `xml:"tilecount,attr"`
	Columns       int            `xml:"columns,attr"`
	ImageMetadata Image          `xml:"image"`
	Tiles         []*TilesetTile `xml:"tile"`
	Texture       *sdl.Texture
}

func (t *Tileset) LoadDataFromTSXFile() (*Tileset, error) {
	// go to tsx folder
	_ = os.Chdir("assets/tilesets/")
	tmxFile, err := os.Open(t.Source)
	// go back to root
	_ = os.Chdir("../../")

	if err != nil {
		return nil, fmt.Errorf("error opening TMX file %s: %v", t.Source, err)
	}

	defer tmxFile.Close()

	bytes, err := ioutil.ReadAll(tmxFile)
	if err != nil {
		return nil, fmt.Errorf("error reading TMX file %s: %v", t.Source, err)
	}

	err = xml.Unmarshal(bytes, t)
	if err != nil {
		return nil, fmt.Errorf("only <xml> format is allowed: %v", err)
	}
	return t, nil
}
func (t *Tileset) LoadTileSetTexture(renderer *sdl.Renderer) {
	_ = os.Chdir("assets/tilesets/")
	surface, err := img.Load(t.ImageMetadata.Source)
	os.Chdir("../../")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		os.Exit(2)
	}
	defer surface.Free()

	t.Texture, err = renderer.CreateTextureFromSurface(surface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(2)
	}
}
func (t *Tileset) GetTileRectSliceFromTilesetByID(id int32) sdl.Rect {
	// Coordinates of sprite slice
	x0 := (int(id) % t.Columns) * t.TileWidth
	y0 := (int(id) / t.Columns) * t.TileHeight
	src := sdl.Rect{int32(x0), int32(y0), int32(t.TileWidth), int32(t.TileHeight)}
	return src
}
