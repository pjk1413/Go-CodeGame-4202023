package players

import (
	"image/color"

	"codegame.com/codegame/tank"
)

type TemplateTanks struct {
	Name  string
	Color color.Color
}

// Will need the tilemap and the player locations
func (p *TemplateTanks) Update(t *tank.Tank) {
	pos := t.GetPosition()
	t.SetFacing(.8)
	if pos.Facing < .9 && pos.Facing > .7 {
		t.Fire()
	} else {
		t.RotateRight()
	}

	t.Update()
}

func NewTank() *tank.Tank {
	t := &tank.Tank{
		Name:  "Template Tank",
		Color: color.Black,
	}

	return t
}
