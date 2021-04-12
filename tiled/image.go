package tiled

import (
	"encoding/xml"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"io/ioutil"
	"os"
)

type Image struct {
	XMLName xml.Name `xml:"image"`
	Format  string   `xml:"format"`
	Source  string   `xml:"source,attr"`
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
	ImageMetadata      Image    `xml:"image"`
	Image *ebiten.Image
}

func (t *Tileset) GetTileImgByID(id int) *ebiten.Image {
	// The tsx format starts counting tiles from 1, so to make these calculations
	// work correctly, we need to decrement the ID by 1
	id -= 1

	x0 := (id % t.Columns) * t.TileWidth
	y0 := (id / t.Columns) * t.TileHeight
	x1, y1 := x0+t.TileWidth, y0+t.TileHeight

	return t.Image.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
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
func (t *Tileset) LoadTileSetImage(){
	var tilesetIMG *ebiten.Image
	{
		// change to tileset dir
		_ = os.Chdir("assets/tilesets/")
		imgFile, err := os.Open(t.ImageMetadata.Source)
		// change back to root
		os.Chdir("../../")
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		img, _, err := image.Decode(imgFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		tilesetIMG = ebiten.NewImageFromImage(img)
	}
	t.Image = tilesetIMG
}