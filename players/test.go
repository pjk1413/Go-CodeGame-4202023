package players

import (
	"image/color"

	"codegame.com/codegame/tank"
)

type Destroyer struct {
	Name  string
	Color color.Color
}

func (p *Destroyer) Update(t *tank.Tank) {
	// t.Decelerate()
	// t.RotateLeft()
	t.Update()
}
