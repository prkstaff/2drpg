package main

import (
	"embed"
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
	Window *sdl.Window
	Surface *sdl.Surface
	windowName string
}

func (g Game) run()  {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	var err error
	g.Window, err = sdl.CreateWindow(g.windowName, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(settings.GameSettings().WindowWidth), int32(settings.GameSettings().WindowHeigh), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer g.Window.Destroy()

	g.Surface, err = g.Window.GetSurface()
	if err != nil {
		panic(err)
	}
	g.Surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 200, 200}
	g.Surface.FillRect(&rect, 0xffff0000)
	g.Window.UpdateSurface()
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
			g.Update()
			g.Draw()
		}
	}
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
func (g *Game) Draw() {
	// Write your game's rendering.

	//g.scenes[g.currentScene].Draw(screen)
}

func main() {
	game := Game{windowName: "Unstopable Fellowship"}

	//start Screen
	villageScene := scenes.VillageScene{}
	villageScene.EmbeddedFS = &embeddedFS
	game.scenes = map[string]scenes.Scene{"village": &villageScene}

	//set scene
	game.loadScene("village")

	game.run()
}
