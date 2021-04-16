package main

import (
	"embed"
	"fmt"
	"github.com/prkstaff/2drpg/scenes"
	"github.com/prkstaff/2drpg/settings"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

// embeddedFS holds our game assets so we can distribute our game as a single binary
//go:embed assets/maps/main.tmx assets/tilesets/tileset.png assets/tilesets/tileset.tsx
var embeddedFS embed.FS

// Game implements ebiten.Game interface.
type Game struct {
	scenes       map[string]scenes.Scene
	currentScene string
	Window *sdl.Window
	windowName string
	renderer *sdl.Renderer
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
	g.renderer , err = sdl.CreateRenderer(g.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
	defer g.renderer.Destroy()

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
func (g *Game) Update() {
	// Write your game's logical update.
	g.scenes[g.currentScene].Update()
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw() {
	var points []sdl.Point
	var rect sdl.Rect
	var rects []sdl.Rect
	renderer := g.renderer
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.DrawPoint(150, 300)

	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.DrawLine(0, 0, 200, 200)

	points = []sdl.Point{{0, 0}, {100, 300}, {100, 300}, {200, 0}}
	renderer.SetDrawColor(255, 255, 0, 255)
	renderer.DrawLines(points)

	rect = sdl.Rect{300, 0, 200, 200}
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawRect(&rect)

	rects = []sdl.Rect{{400, 400, 100, 100}, {550, 350, 200, 200}}
	renderer.SetDrawColor(0, 255, 255, 255)
	renderer.DrawRects(rects)

	rect = sdl.Rect{250, 250, 200, 200}
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.FillRect(&rect)

	rects = []sdl.Rect{{500, 300, 100, 100}, {200, 300, 200, 200}}
	renderer.SetDrawColor(255, 0, 255, 255)
	renderer.FillRects(rects)

	renderer.Present()
	sdl.Delay(16)
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
