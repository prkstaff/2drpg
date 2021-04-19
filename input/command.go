package input

import "github.com/prkstaff/2drpg/characters"

type Command interface {
	Execute(actor characters.Actor)
}

type InputHandler struct {
	commands []Command
}

func (h InputHandler) HandleInput(actor characters.Actor)  {
	if len(h.commands) > 0 {
		for _, command := range h.commands{
			command.Execute(actor)
		}
	}
}