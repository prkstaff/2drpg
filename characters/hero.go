package characters

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/prkstaff/2drpg/settings"
	"github.com/prkstaff/2drpg/tiled"
	"image"
	"os"
)

type Hero struct{
	SpritePath string
	Sprite *ebiten.Image
	SpriteWidth uint16
	SpriteHeight uint16
	XPos uint16
	YPos uint16
	DrawAfterLayer uint8
}

func (h *Hero) LoadSpriteIMG()  {
	var spriteImage *ebiten.Image
	{
		imgFile, err := os.Open(h.SpritePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		img, _, err := image.Decode(imgFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		spriteImage = ebiten.NewImageFromImage(img)
	}
	h.Sprite = spriteImage
}

func (h Hero) Draw(screen *ebiten.Image, gameMap tiled.Map) {
	op := &ebiten.DrawImageOptions{}
	sx := float64(settings.GameSettings().ScreenWidth / (gameMap.Width * gameMap.TileWidth))
	sy := float64(settings.GameSettings().ScreenHeight / (gameMap.Height * gameMap.TileHeight))
	i := 1
	op.GeoM.Translate(
		float64((uint16(i)%gameMap.Width)*gameMap.TileWidth),
		float64((uint16(i)/gameMap.Height)*gameMap.TileHeight),
	)
	op.GeoM.Scale(sx, sy)

	screen.DrawImage(h.GetSpriteImgByID(int(1)), op)
}

func (h *Hero) GetSpriteImgByID(id int) *ebiten.Image {
	// The tsx format starts counting tiles from 1, so to make these calculations
	// work correctly, we need to decrement the ID by 1
	id -= 1

	//x0 := (id % t.Columns) * t.TileWidth
	//y0 := (id / t.Columns) * t.TileHeight
	//x1, y1 := x0+t.TileWidth, y0+t.TileHeight

	return h.Sprite.SubImage(image.Rect(0, 0, 16, 32)).(*ebiten.Image)
}