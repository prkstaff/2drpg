package main

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/prkstaff/2drpg/scenes"
	"github.com/prkstaff/2drpg/settings"
	"github.com/veandco/go-sdl2/sdl"
)

// embeddedFS holds our game assets so we can distribute our game as a single binary
//go:embed assets/maps/main.tmx assets/tilesets/tileset.png assets/tilesets/tileset.tsx
var embeddedFS embed.FS

// Game implements ebiten.Game interface.
type Game struct {
	scenes       map[string]scenes.Scene
	currentScene string
}

func (g *Game) loadScene(sceneName string) {
	g.currentScene = sceneName
	g.scenes[g.currentScene].OnLoad()
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	g.scenes[g.currentScene].Update()
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.

	g.scenes[g.currentScene].Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(settings.GameSettings().LayoutWidth), int(settings.GameSettings().LayoutHeight)
}

func main() {
	//game := &Game{}

	////start Screen
	//villageScene := scenes.VillageScene{}
	//villageScene.EmbeddedFS = &embeddedFS
	//game.scenes = map[string]scenes.Scene{"village": &villageScene}

	////set scene
	//game.loadScene("village")

	// Specify the window size as you like. Here, a doubled size is specified.
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}
