package tiled

import (
	"encoding/xml"
	"fmt"
	"github.com/prkstaff/2drpg/settings"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
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
	Texture *sdl.Texture
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
func (t *Tileset) LoadTileSetTexture(renderer *sdl.Renderer){
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
	// to-do tentar desenhar aqui a img inteira
	defer t.Texture.Destroy()
	src := sdl.Rect{0, 0, 16, 16}
	dst := sdl.Rect{0, 0, int32(settings.GameSettings().WindowWidth), int32(settings.GameSettings().WindowHeigh)}
	renderer.Clear()
	renderer.Copy(t.Texture, &src, &dst)
	renderer.Present()
	sdl.Delay(10000)
}