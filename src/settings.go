package src

import "sync"

// type global
type gameSettings struct {
	ScreenWidth  uint16 // O tamanho do tileset total
	ScreenHeight uint16
	LayoutWidth  uint16 // O tamanho do Layout que vai renderizar
	LayoutHeight uint16
	WindowWidth  uint16
	WindowHeigh  uint16
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
			WindowHeigh:  800,
		} // <-- thread safe

	})
	return instance
}
