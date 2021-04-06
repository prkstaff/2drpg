package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface.
type Game struct {
	scenes       map[string]Scene
	currentScene string
}

func (g *Game) loadScene(sceneName string) {
	g.currentScene = sceneName
	g.scenes[g.currentScene].onLoad()
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	g.scenes[g.currentScene].update()
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.

	g.scenes[g.currentScene].draw(screen)
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
	startScreen := GameStartScreen{}
	game.scenes = map[string]Scene{"start": &startScreen}

	//set scene
	game.loadScene("start")

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(1024, 1000)
	ebiten.SetWindowTitle("My 2D RPG")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
