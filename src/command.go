package src

type Command interface {
	Execute(actor Actor)
}

type InputHandler struct {
	Commands []Command
}

func (h InputHandler) HandleInput(actor Actor) {
	if len(h.Commands) > 0 {
		for _, command := range h.Commands {
			command.Execute(actor)
		}
	}
}
