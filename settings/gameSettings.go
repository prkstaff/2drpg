package settings

import "sync"

// type global
type gameSettings struct {
	ScreenWidth  int // O tamanho do tileset total
	ScreenHeight int
	LayoutWidth  int // O tamanho do Layout que vai renderizar
	LayoutHeight int
	WindowWidth  int
	WindowHeigh  int
}

var instance *gameSettings = nil
var once sync.Once

func GameSettings() *gameSettings {
	once.Do(func() {
		instance = &gameSettings{
			ScreenWidth:  800,
			ScreenHeight: 800,
			LayoutWidth:  400,
			LayoutHeight: 400,
			WindowWidth:  1024,
			WindowHeigh:  1000,
		} // <-- thread safe

	})
	return instance
}
