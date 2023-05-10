package players

import (
	"image/color"

	"codegame.com/codegame/tank"
)

type BigBoy struct{}

func (d *BigBoy) Update(t *tank.Tank, pos []tank.TankPosition) {
	t.Accelerate()
	t.Update()
}

func (d *BigBoy) Create() *tank.Tank {
	// Create a new tank
	return &tank.Tank{
		Name:  "BigBoy",
		Color: color.RGBA{0xff, 0, 0xff, 0xff},
	}
}
