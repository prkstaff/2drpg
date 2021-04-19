package input

import "github.com/prkstaff/2drpg/characters"

type MoveUpCommand struct{}

func (mu MoveUpCommand) Execute(actor characters.Actor) {
	actor.MoveUp()
}

type MoveDownCommand struct{}

func (md MoveDownCommand) Execute(actor characters.Actor) {
	actor.MoveDown()
}

type MoveLeftCommand struct{}

func (ml MoveLeftCommand) Execute(actor characters.Actor) {
	actor.MoveLeft()
}

type MoveRightCommand struct{}

func (mr MoveRightCommand) Execute(actor characters.Actor) {
	actor.MoveRight()
}
