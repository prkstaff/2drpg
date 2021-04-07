package main

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/prkstaff/2drpg/scenes"
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

const (
	width  = 400
	height = 400
)

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	game := &Game{}

	//start Screen
	villageScene := scenes.VillageScene{}
	villageScene.EmbeddedFS = &embeddedFS
	game.scenes = map[string]scenes.Scene{"village": &villageScene}

	//set scene
	game.loadScene("village")

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(1024, 1000)
	ebiten.SetWindowTitle("My 2D RPG")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
