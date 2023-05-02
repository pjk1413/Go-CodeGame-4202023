package players

import (
	"image/color"
	"math"

	"codegame.com/codegame/tank"
)

type Destroyer struct {
	Name  string
	Color color.Color
}

func (d *Destroyer) Update(t *tank.Tank, pos []tank.TankPosition) {
	p := t.GetPosition()

	degrees := p.Facing * (180 / math.Pi)
	if degrees < -85 && degrees > -95 {
		t.Accelerate()
	} else {
		t.RotateLeft()
	}

	t.Update()
}

func (d *Destroyer) Create() *tank.Tank {
	// Create a new tank
	return &tank.Tank{
		Name:  "Destroyer",
		Color: color.RGBA{0, 0, 0xff, 0xff},
	}
}
