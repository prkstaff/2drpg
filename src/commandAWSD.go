package src

type MoveUpCommand struct{}

func (mu MoveUpCommand) Execute(actor Actor) {
	actor.MoveUp()
}

type MoveDownCommand struct{}

func (md MoveDownCommand) Execute(actor Actor) {
	actor.MoveDown()
}

type MoveLeftCommand struct{}

func (ml MoveLeftCommand) Execute(actor Actor) {
	actor.MoveLeft()
}

type MoveRightCommand struct{}

func (mr MoveRightCommand) Execute(actor Actor) {
	actor.MoveRight()
}
