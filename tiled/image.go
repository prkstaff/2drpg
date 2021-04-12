package tiled

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Image struct {
	XMLName xml.Name `xml:"image"`
	Format  string   `xml:"format"`
	Source  string   `xml:"source"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
}

type Tileset struct {
	Source string `xml:"source,attr"`
	XMLName    xml.Name `xml:"tileset"`
	Version    string   `xml:"version,attr"`
	Name       string   `xml:"name,attr"`
	TileWidth  int      `xml:"tilewidth,attr"`
	TileHeight int      `xml:"tileheight,attr"`
	TileCount  int      `xml:"tilecount,attr"`
	Columns    int      `xml:"columns,attr"`
	Image      Image    `xml:"image"`
}

func (t *Tileset) LoadDataFromTSXFile() (*Tileset, error){
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