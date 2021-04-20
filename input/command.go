package input

import "github.com/prkstaff/2drpg/characters"

type Command interface {
	Execute(actor characters.Actor)
}

type InputHandler struct {
	Commands []Command
}

func (h InputHandler) HandleInput(actor characters.Actor) {
	if len(h.Commands) > 0 {
		for _, command := range h.Commands {
			command.Execute(actor)
		}
	}
}
