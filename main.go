package main

import (
	"embed"
	"fmt"
	"os"

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
	Window       *sdl.Window
	windowName   string
	renderer     *sdl.Renderer
	keyStates    []uint8
}

func (g Game) run() {
	g.keyStates = sdl.GetKeyboardState()

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
	g.renderer, err = sdl.CreateRenderer(g.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
	defer g.renderer.Destroy()

	g.loadScene("village")

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
	g.scenes[g.currentScene].OnLoad(g.renderer)
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() {
	// Write your game's logical update.
	g.scenes[g.currentScene].Update(g.keyStates)
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw() {
	// fazer um blog post
	// https://stackoverflow.com/questions/21007329/what-is-an-sdl-renderer

	g.renderer.Clear()
	g.scenes[g.currentScene].Draw(g.renderer)
	g.renderer.Present()
	sdl.Delay(16)
}

func main() {
	game := Game{windowName: "Unstopable Fellowship"}

	//start Screen
	villageScene := scenes.VillageScene{}
	villageScene.EmbeddedFS = &embeddedFS
	game.scenes = map[string]scenes.Scene{"village": &villageScene}

	game.run()
}
