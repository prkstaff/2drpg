package tiled

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Data struct {
	// The encoding used to encode the tile layer data. When used, it can be "base64" and "csv" at the moment.
	Encoding string `xml:"encoding,attr"`
	// The compression used to compress the tile layer data. Tiled Qt supports "gzip" and "zlib".
	Compression string `xml:"compression,attr"`
	// Raw data
	RawData []byte `xml:",innerxml"`
	// RawData string `xml:",innerxml"`
}
type Layer struct {
	Width float64 `xml:"width,attr"`
	Height float64 `xml:"height,attr"`
	ID uint32 `xml:"id,attr"`
	Name string `xml:"name,attr"`
	Data Data`xml:"data"`
}

type Object struct {
	// Unique ID of the object. Each object that is placed on a map gets a unique id. Even if an object was deleted, no object gets the same ID.
	// Can not be changed in Tiled Qt. (since Tiled 0.11)
	ID uint32 `xml:"id,attr"`
	// The name of the object. An arbitrary string.
	Name string `xml:"name,attr"`
	// The type of the object. An arbitrary string.
	Type string `xml:"type,attr"`
	// The x coordinate of the object.
	X float64 `xml:"x,attr"`
	// The y coordinate of the object.
	Y float64 `xml:"y,attr"`
	// The width of the object (defaults to 0).
	Width float64 `xml:"width,attr"`
	// The height of the object (defaults to 0).
	Height float64 `xml:"height,attr"`
}

// ObjectGroup is in fact a map layer, and is hence called "object layer" in Tiled Qt
type ObjectGroup struct {
	// Unique ID of the layer.
	// Each layer that added to a map gets a unique id. Even if a layer is deleted,
	// no layer ever gets the same ID. Can not be changed in Tiled. (since Tiled 1.2)
	ID uint32 `xml:"id,attr"`
	// The name of the object group.
	Name string `xml:"name,attr"`
	// Group objects
	Objects []*Object `xml:"object"`
}

type Map struct{
	Width uint16 `xml:"width,attr"` // Number of tiles Width
	Height uint16 `xml:"height,attr"` // Number of tiles Height
	TileWidth uint16 `xml:"tilewidth,attr"` // Tile Width size in pixels
	TileHeight uint16 `xml:"tileheight,attr"` // Tile Height size in pixels
	Infinite     bool     `xml:"infinite,attr"`
	RenderOrder  string   `xml:"renderorder,attr"`
	Layers []*Layer `xml:"layer"`
	Source string `xml:"source,attr"` // The tsx file
	Version      string   `xml:"version,attr"`
	TiledVersion string   `xml:"tiledversion,attr"`
	Objects ObjectGroup `xml:"objectgroup"`
	Orientation  string   `xml:"orientation,attr"`
}

func ReadTMX(tmxPath string){
	data, err := ioutil.ReadFile(tmxPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	var tiledMap Map
	xml.Unmarshal(data, &tiledMap)
	fmt.Println(tiledMap)
}
