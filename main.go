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

	g.loadScene("village")

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type)  {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				keyCode := t.Keysym.Sym
				keys := ""

				// Modifier keys
				switch t.Keysym.Mod {
				case sdl.KMOD_LALT:
					keys += "Left Alt"
				case sdl.KMOD_LCTRL:
					keys += "Left Control"
				case sdl.KMOD_LSHIFT:
					keys += "Left Shift"
				case sdl.KMOD_LGUI:
					keys += "Left Meta or Windows key"
				case sdl.KMOD_RALT:
					keys += "Right Alt"
				case sdl.KMOD_RCTRL:
					keys += "Right Control"
				case sdl.KMOD_RSHIFT:
					keys += "Right Shift"
				case sdl.KMOD_RGUI:
					keys += "Right Meta or Windows key"
				case sdl.KMOD_NUM:
					keys += "Num Lock"
				case sdl.KMOD_CAPS:
					keys += "Caps Lock"
				case sdl.KMOD_MODE:
					keys += "AltGr Key"
				}

				if keyCode < 10000 {
					if keys != "" {
						keys += " + "
					}

					// If the key is held down, this will fire
					if t.Repeat > 0 {
						keys += string(keyCode) + " repeating"
					} else {
						if t.State == sdl.RELEASED {
							keys += string(keyCode) + " released"
						} else if t.State == sdl.PRESSED {
							keys += string(keyCode) + " pressed"
						}
					}

				}

				if keys != "" {
					fmt.Println(keys)
				}
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
	g.scenes[g.currentScene].Update()
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
