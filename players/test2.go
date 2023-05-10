package players

import (
	"image/color"

	"codegame.com/codegame/tank"
)

type Killer struct{}

func (d *Killer) Update(t *tank.Tank, pos []tank.TankPosition) {
	t.Decelerate()
	t.RotateLeft()
	t.Update()
}

func (d *Killer) Create() *tank.Tank {
	// Create a new tank
	return &tank.Tank{
		Name:  "Killer",
		Color: color.RGBA{0, 0xff, 0xff, 0xff},
	}
}
